package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/lib/pq"
	"log"
	"os"
)

var db sql.DB
var dbmap *gorp.DbMap

type AMP struct {
	Id   int64
	Apid string `xml:"APID"`
	Vpid string `xml:"VPID"`
	Name string `xml:"NM"`
}

// Initialises the connection to the database (based on configuration)
// and sets up the ORM ready for use.
func db_init() error {
	connstr, err := pq.ParseURL(config.Database)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid connection string: %s", config.Database))
	}

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return err
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	if config.Debug {
		dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	}

	dbmap.AddTableWithName(AMP{}, "actual_medical_product").SetKeys(true, "Id")
	dbmap.CreateTablesIfNotExists()

	return nil
}

// Clean up the DB, basically just drop the tables
func db_clean() {
	dbmap.DropTables()
}

// Close the conenction to the database
func db_close() {
	if config.Debug {
		dbmap.TraceOff()
	}
	db.Close()
}
