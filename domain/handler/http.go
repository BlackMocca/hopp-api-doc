package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type HttpHandler struct {
	datasource domain.Datasource
}

func NewHttpHandler(datasource domain.Datasource) HttpHandler {
	return HttpHandler{datasource: datasource}
}

func (h HttpHandler) readMetadata(path string) (map[string]interface{}, error) {
	var meta = make(map[string]interface{})
	f, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return meta, err
	}
	if errors.Is(err, os.ErrNotExist) {
		return meta, nil
	}

	if err := json.Unmarshal(f, &meta); err != nil {
		return meta, err
	}

	return meta, nil
}

func (h HttpHandler) fillTeamCollectionMetaData(teams []models.Team) error {
	if len(teams) == 0 {
		return nil
	}
	meta, err := h.readMetadata(constants.TEAM_META_DATA_PATH)
	if err != nil {
		return err
	}

	if len(meta) > 0 {
		for index := range teams {
			if v, ok := meta[teams[index].Id]; ok {
				ti, err := time.Parse(time.RFC3339, cast.ToString(v.(map[string]interface{})["last_updated"]))
				if err != nil {
					return err
				}
				teams[index].LastUpdated = ti.Format(constants.TIMESTAMP_LAYOUT)
			}
		}
	}

	return nil
}

func (h HttpHandler) Index(c echo.Context) error {
	var ctx = context.Background()
	fmt.Println("cookie context", c.Cookies())

	teams, err := h.datasource.FetchAllTeams(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := h.fillTeamCollectionMetaData(teams); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"teams": teams,
	}
	return c.Render(http.StatusOK, "index", resp)
}
