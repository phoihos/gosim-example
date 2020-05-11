package main

import (
	"log"

	"github.com/phoihos/gosim/database"
	"github.com/phoihos/gosim/database/mssql"
	"github.com/phoihos/gosim/server"

	_ "handler"
)

func main() {
	defer database.Close()

	dbOpt := &database.SetupOption{Alias: "example", Host: "127.0.0.1", Port: "1433", Database: "exam", User: "user", Password: "password"}
	dbConf := mssql.NewConfiguration(dbOpt)
	if err := database.OpenConnection(dbConf); err != nil {
		log.Print(err)
	}

	conf := &server.Configuration{Port: "8080", ShutdownPath: "/shutdown"}
	server.Run(conf)
}
