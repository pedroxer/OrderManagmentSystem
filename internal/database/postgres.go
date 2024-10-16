package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pedroxer/ordermanagmentsystem/internal/config"
	"github.com/pedroxer/ordermanagmentsystem/internal/storage"
)

func ConnectToPostgres(cfg config.Postgres) (storage.DBTX, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Db,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
