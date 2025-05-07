package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func RpcHealthCheck(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	// This is a health check function that can be called from the client.
	// It returns a simple "OK" message to indicate that the server is running.
	return "OK", nil
}
