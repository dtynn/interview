package db

import (
	"context"

	"github.com/dtynn/interview/proto"
	"github.com/jmoiron/sqlx"
)

var Team team

type team struct {
}

func (t *team) List(ctx context.Context, limit int) ([]proto.Team, error) {
	teams := make([]proto.Team, 0, 32)
	if err := Invoke(func(db *sqlx.DB) error {
		query := `
		SELECT
			id,
			name
		FROM interview.team
		ORDER BY id ASC
		LIMIT $1
		`

		return db.SelectContext(ctx, &teams, query, limit)
	}); err != nil {
		return nil, err
	}

	return teams, nil
}

func (t *team) Add(ctx context.Context, teams ...proto.Team) error {
	return InvokeTx(func(tx *sqlx.Tx) error {
		stmt, err := tx.Prepare(`
			INSERT INTO interview.team (
				id,
				name
			) VALUES (
				$1,
				$2
			)
		`)

		if err != nil {
			return err
		}

		for _, team := range teams {
			_, err := stmt.ExecContext(ctx, team.ID, team.Name)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
