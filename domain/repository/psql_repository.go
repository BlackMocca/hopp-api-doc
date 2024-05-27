package repository

import (
	"context"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type psqlRepository struct {
	db *pgxpool.Pool
}

func NewPsqlRepository(client *pgxpool.Pool) domain.Datasource {
	return &psqlRepository{db: client}
}

func (p psqlRepository) FetchAllTeams(ctx context.Context) ([]models.Team, error) {
	sql := `SELECT * FROM "Team"`
	rows, err := p.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Team])
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (p psqlRepository) FetchAllTeamCollection(ctx context.Context, teamId string) ([]models.TeamCollection, error) {
	sql := `
	SELECT t1.id, t1."parentID", t1."teamID", t1."title", t1."data"
	FROM "TeamCollection" t1
	LEFT JOIN "TeamCollection" t2
	ON t1."parentID" = t2.id
	WHERE t1."teamID" = $1::text
	ORDER BY t1."parentID" ASC, t1."orderIndex" ASC 
	`
	rows, err := p.db.Query(ctx, sql, teamId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ptrs, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.TeamCollection])
	if err != nil {
		return nil, err
	}

	return ptrs, nil
}

func (p psqlRepository) FetchAllTeamRequest(ctx context.Context, teamId string) ([]models.TeamRequest, error) {
	sql := `
	SELECT id, "collectionID", "teamID", title, request
	FROM "TeamRequest"
	where "teamID" = $1::text
	ORDER BY "collectionID" ASC, "orderIndex" ASC
	`
	rows, err := p.db.Query(ctx, sql, teamId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ptrs, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.TeamRequest])
	if err != nil {
		return nil, err
	}

	return ptrs, nil
}

func (p psqlRepository) fetchInfraConfigs(ctx context.Context, pattern string) ([]models.InfraConfig, error) {
	sql := `SELECT id, name, value, active FROM "InfraConfig" WHERE name LIKE $1::text`
	rows, err := p.db.Query(ctx, sql, pattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ptrs, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.InfraConfig])
	if err != nil {
		return nil, err
	}

	return ptrs, nil
}

func (p psqlRepository) FetchAuthProviders(ctx context.Context) ([]string, error) {
	var providers = make([]string, 0)
	configs, err := p.fetchInfraConfigs(ctx, "VITE_ALLOWED_AUTH_PROVIDERS")
	if err != nil {
		return nil, err
	}
	if len(configs) == 0 {
		return providers, nil
	}
	for _, item := range configs {
		providers = append(providers, item.Value)
	}

	return providers, nil
}

func (p psqlRepository) FetchOneUser(ctx context.Context, email string) (*models.User, error) {
	sql := `SELECT uid, "displayName", email, "photoURL", "isAdmin" FROM "User" WHERE email = $1::text`
	rows, err := p.db.Query(ctx, sql, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ptr, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.User])
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	if ptr.Id == "" {
		return nil, nil
	}
	return &ptr, nil
}
