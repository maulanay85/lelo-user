package role

import (
	"context"
	entity "lelo-user/entity"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func (r *RoleRepositoryModule) GetRoleByCode(ctx context.Context, tx pgx.Tx, code string) (*entity.RoleEntity, error) {
	var role entity.RoleEntity

	err := tx.QueryRow(ctx,
		`SELECT 
			id,
			name,
			code,
			status
		FROM
			t_mst_role
		WHERE
			status = 1
			AND code = $1
		`, code,
	).Scan(&role.Id, &role.Name, &role.Code, &role.Status)
	if err != nil {
		log.Errorf("[repository]: GetRoleByCode err: %v", err)
		return nil, err
	}
	return &role, nil
}
