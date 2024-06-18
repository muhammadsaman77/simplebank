package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhammadsaman77/simplebank/api"
	db "github.com/muhammadsaman77/simplebank/db/sqlc"
	"github.com/muhammadsaman77/simplebank/util"
)

var DB *pgxpool.Pool



func main() {
	conf,err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(conf.DBSource)
	if err != nil {
		log.Fatal("error config db:", err)
	}
	config.MaxConns = 20
	config.MaxConnIdleTime = time.Minute
	DB, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer DB.Close()
	store:= db.NewStore(DB)
	server,_:= api.NewServer(conf, store)
	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}