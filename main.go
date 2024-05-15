package main

import (
	"context"
	"database/sql"
	_ "database/sql"
	"log"
	"net/http"

	"github.com/adotomov/sf-sequence-api/internal"
	"github.com/adotomov/sf-sequence-api/internal/config"
	"github.com/adotomov/sf-sequence-api/internal/handlers"
	db "github.com/adotomov/sf-sequence-api/internal/repository/sql"
	"github.com/adotomov/sf-sequence-api/internal/service/sequence"
	"github.com/adotomov/sf-sequence-api/internal/service/step"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	var c config.AppConfig
	ctx := context.Background()
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	// Connect to database
	sqldb, err := sql.Open(c.DBConfig.Driver, string(c.DBConfig.DSN))
	if err != nil {
		log.Fatal(err)
	}
	sqldb.SetMaxIdleConns(c.DBConfig.MaxIdleConns)
	sqldb.SetMaxOpenConns(c.DBConfig.MaxOpenConns)
	sqldb.SetConnMaxIdleTime(c.DBConfig.ConnMaxIdleTime)
	sqldb.SetConnMaxLifetime(c.DBConfig.ConnMaxLifetime)

	d, err := db.NewDB(sqldb)
	if err != nil {
		log.Fatal(err)
	}

	// Initiate services
	sqsrv := sequence.NewService(c.Env, d)
	stsrv := step.NewService(c.Env, d)
	log.Print("Service initiated")

	// Stand up Http server
	sh := handlers.NewSequenceHandler(sqsrv)
	sth := handlers.NewStepHandler(stsrv)
	router := internal.NewRouter(sh, sth)

	log.Fatal(http.ListenAndServe(":3000", router))
}
