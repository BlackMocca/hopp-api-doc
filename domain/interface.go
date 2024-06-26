package domain

import (
	"context"

	"github.com/Blackmocca/hopp-api-doc/domain/models"
)

type Datasource interface {
	FetchAllTeams(ctx context.Context) ([]models.Team, error)
	FetchAllTeamCollection(ctx context.Context, teamId string) ([]models.TeamCollection, error)
	FetchAllTeamRequest(ctx context.Context, teamId string) ([]models.TeamRequest, error)
	FetchAuthProviders(ctx context.Context) ([]string, error)
	FetchOneUser(ctx context.Context, email string) (*models.User, error)

	// fetch data with user id
	FetchMyTeams(ctx context.Context, userId string) ([]models.Team, error)
}

type AuthProvider interface {
	LinkAuth(ctx context.Context) (string, error)
	SignToken(ctx context.Context, authorizeCode string, state string) (*models.JWTToken, error)
	Me(ctx context.Context, jwtToken string) (*models.ProfileProvider, error)
}
