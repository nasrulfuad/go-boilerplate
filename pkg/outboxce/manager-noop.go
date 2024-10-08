package outboxce

import (
	"context"
	"database/sql"
)

var _ Manager = managerNOP{}

type managerNOP struct{}

func (n managerNOP) Store(ctx context.Context, tx *sql.Tx, ob OutboxCE) (err error) {
	return
}

func (n managerNOP) RelayLoop(ctx context.Context, relay RelayFunc) (err error) {
	<-ctx.Done()
	return ctx.Err()
}

func ManagerNOP() Manager { return managerNOP{} }
