package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testQueries *Queries
const (
	dbSource = "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable"
)
func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err:= pgx.Connect(ctx, dbSource);
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close(ctx)

	testQueries = New(conn);

	os.Exit(m.Run())

}