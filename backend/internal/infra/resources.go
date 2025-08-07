package infra

import (
	"context"
	"fmt"

	"booklib/internal/infra/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rizanw/go-log"
)

type Resources struct {
	// rdbms
	Database *sqlx.DB
}

func NewResources(ctx context.Context, conf *config.Config) (*Resources, error) {
	// set log config
	err := log.SetConfig(
		&log.Config{
			AppName:     conf.AppName,
			Environment: conf.Env,
		},
	)
	if err != nil {
		log.Fatal(ctx, err, nil, "failed to set log config")
	}

	// init db
	psqlsource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.DBName)
	psqlx, err := sqlx.Connect("postgres", psqlsource)
	if err != nil {
		log.Fatalf(ctx, err, nil, "Failed create database conn, with err:%+v", err)
		return nil, err
	}

	return &Resources{
		Database: psqlx,
	}, nil
}
