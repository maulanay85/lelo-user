package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (u *UtilDBModule) BeginTx(ctx context.Context, opt ...pgx.TxOptions) (pgx.Tx, error) {
	var tx pgx.Tx
	var err error

	if len(opt) > 0 {
		tx, err = u.Db.BeginTx(ctx, opt[0])
	} else {
		tx, err = u.Db.BeginTx(ctx, pgx.TxOptions{})
	}

	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	return tx, err
}
