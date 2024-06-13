package db

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool
const (
	dbSource = "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable"
)
func TestMain(m *testing.M) {
	ctx := context.Background()
	var err error
	config, err := pgxpool.ParseConfig(dbSource) 
	if err != nil {
		log.Fatal("error config db:", err)
	}
	config.MaxConns = 20
	config.MaxConnIdleTime = time.Minute
	testDB, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer testDB.Close()

	testQueries = New(testDB);

	os.Exit(m.Run())

}