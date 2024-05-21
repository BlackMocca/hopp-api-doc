package domain

import (
	"context"

	"github.com/Blackmocca/hopp-api-doc/domain/models"
)

type Datasource interface {
	FetchAllTeams(ctx context.Context) ([]models.Team, error)
	FetchAllTeamCollection(ctx context.Context, teamId string) ([]models.TeamCollection, error)
	FetchAllTeamRequest(ctx context.Context, teamId string) ([]models.TeamRequest, error)
}
