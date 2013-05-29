package main

import (
	"database/sql"
	_ "fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// The database connection
var db sql.DB
var dbmap *gorp.DbMap

type Invoice struct {
	Id       int64
	Created  int64
	Updated  int64
	Memo     string
	PersonId int64
}

func db_init() error {
	db, err := sql.Open("postgres", config.Database)
	if err != nil {
		return err
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	dbmap.AddTableWithName(Invoice{}, "invoice_test").SetKeys(true, "Id")
	dbmap.CreateTablesIfNotExists()

	return nil
}

func db_close() {
	dbmap.TraceOff()
	db.Close()
}
