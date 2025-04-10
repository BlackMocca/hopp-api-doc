package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	docs "github.com/Blackmocca/hopp-api-doc/hopp-cli/methods"
	"github.com/go-co-op/gocron/v2"
	"github.com/gosimple/slug"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cast"
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
	ti := cast.ToInt(constants.TIMER_SYNC_TEAM_COLLECTION)
	c.cron.NewJob(gocron.DurationJob(time.Duration(ti)*time.Second), gocron.NewTask(func() {
		var ctx = context.Background()
		var teams, err = c.repository.FetchAllTeams(ctx)
		if err != nil {
			log.Errorf("failed to fetch all teams: %s", err.Error())
			return
		}
		if len(teams) == 0 {
			return
		}

		var teamMetadata = make(map[string]interface{}, 0)
		for _, team := range teams {
			var dirTeam = filepath.Join("./assets", team.Id)
			os.Mkdir(dirTeam, 0755)

			colls, err := c.getExportCollection(ctx, team)
			if err != nil {
				log.Errorf("failed to get export collection: %s", err.Error())
				continue
			}
			if len(colls) == 0 {
				continue
			}

			/* process export into docs */
			var filePath = filepath.Join(dirTeam, "./data.json")
			var fileData, _ = json.Marshal(colls)
			var pathOut = fmt.Sprintf("./docs/%s", slug.Make(team.Id))
			os.WriteFile(filePath, fileData, 0644)

			if err := docs.GenerateDocs(pathOut, filePath, 0, false, "./hopp-cli/templates"); err != nil {
				log.Errorf("failed to get export document: %s", err.Error())
				continue
			}

			teamMetadata[team.Id] = map[string]interface{}{
				"id":           team.Id,
				"name":         team.Name,
				"last_updated": time.Now().Format(time.RFC3339),
			}
		}
		teamMetaData, _ := json.Marshal(teamMetadata)
		os.WriteFile(constants.TEAM_META_DATA_PATH, teamMetaData, 0755)

	}))

	c.cron.Start()
}

func (c *CronjobHandler) Shutdown() {
	c.cron.Shutdown()
}
