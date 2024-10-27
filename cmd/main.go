package main

import (
	"encoding/json"
	"github.com/caarlos0/env/v6"
	"github.com/pedroxer/ordermanagmentsystem/internal/app"
	"github.com/pedroxer/ordermanagmentsystem/internal/config"
	"github.com/pedroxer/ordermanagmentsystem/internal/database"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	log := log.New()
	data, err := os.ReadFile("./configs/config.json")
	if err != nil {
		log.Fatal(err)
	}
	cfg := new(config.Config)

	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatal(err)
	}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	conn, err := database.ConnectToPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed connect to db %s", err)
	}

	tokenTTL, err := time.ParseDuration(cfg.TokenTTL)
	if err != nil {
		log.Fatal(err)
	}
	application := app.NewApp(log, cfg.Serv.Port, conn, tokenTTL)

	if err := application.GRPCSrvr.Run(); err != nil {
		log.Fatal(err)
	}
}
