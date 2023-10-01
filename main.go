package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/walteranderson/twtr/api"
	"github.com/walteranderson/twtr/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "The server address")
	dbDriver := flag.String("dbdriver", "sqlite3", "The DB driver")
	dbSourceName := flag.String("dbsourcename", "sqlite.db", "The DB source name")
	flag.Parse()

	db, err := sql.Open(*dbDriver, *dbSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := storage.NewSQLiteStorage(db)
	if err = store.Migrate(); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(*listenAddr, store)
	server.Start()
}
