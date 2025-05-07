package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Debug("This is a Nakama Plugin.")

	if err := initializer.RegisterRpc("healthcheck", RpcHealthCheck); err != nil {
		return err
	}

	return nil
}
