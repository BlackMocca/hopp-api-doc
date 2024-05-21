package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	docs "github.com/Blackmocca/hopp-api-doc/hopp-cli/methods"
	"github.com/go-co-op/gocron/v2"
	"github.com/gosimple/slug"
	"github.com/labstack/gommon/log"
)

type CronjobHandler struct {
	cron       gocron.Scheduler
	repository domain.Datasource
}

func NewCronjobHandler(cron gocron.Scheduler, repository domain.Datasource) CronjobHandler {

	return CronjobHandler{cron: cron, repository: repository}
}

func (c *CronjobHandler) getExportCollection(ctx context.Context, team models.Team) ([]models.Collection, error) {
	tcolls, err := c.repository.FetchAllTeamCollection(ctx, team.Id)
	if err != nil {
		return nil, err
	}
	if len(tcolls) == 0 {
		return nil, nil
	}

	treqs, err := c.repository.FetchAllTeamRequest(ctx, team.Id)
	if err != nil {
		return nil, err
	}

	return c.findTeamCollStack(team, tcolls, treqs, ""), nil
}

func (c *CronjobHandler) findTeamCollStack(team models.Team, tcolls []models.TeamCollection, treqs []models.TeamRequest, parentId string) []models.Collection {
	var filterTColls = models.TeamCollections(tcolls).FindByParentId(parentId)
	if len(filterTColls) == 0 {
		return make([]models.Collection, 0)
	}
	var exportColls = make([]models.Collection, 0)
	for _, item := range filterTColls {
		item.Folders = c.findTeamCollStack(team, tcolls, treqs, item.Id)
		item.Requests = models.TeamRequests(treqs).FindByCollectionId(item.Id)

		exportColls = append(exportColls, models.NewDefaultCollection(team, item))
	}

	return exportColls
}

func (c *CronjobHandler) Run() {
	// c.cron.NewJob(gocron.DurationJob(1*time.Minute), gocron.NewTask(func() {
	log.Info(" Starting Job")
	var ctx = context.Background()
	var teams, err = c.repository.FetchAllTeams(ctx)
	if err != nil {
		log.Errorf("failed to fetch all teams: %s", err.Error())
		return
	}
	if len(teams) == 0 {
		return
	}

	for _, team := range teams {
		var dirTeam = filepath.Join("./assets", team.Id)
		os.Mkdir(dirTeam, 0755)

		colls, err := c.getExportCollection(ctx, team)
		if err != nil {
			log.Errorf("failed to get export collection: %s", err.Error())
			return
		}
		if len(colls) == 0 {
			return
		}

		/* process export into docs */
		var filePath = filepath.Join(dirTeam, "./data.json")
		var fileData, _ = json.Marshal(colls)
		var pathOut = fmt.Sprintf("./docs/%s", slug.Make(team.Id))
		os.WriteFile(filePath, fileData, 0644)

		if err := docs.GenerateDocs(pathOut, filePath, 0, false, "./hopp-cli/templates"); err != nil {
			log.Errorf("failed to get export document: %s", err.Error())
			return
		}
	}

	log.Info("Ending Job")
	// }))

	c.cron.Start()
}

func (c *CronjobHandler) Shutdown() {
	c.cron.Shutdown()
}
