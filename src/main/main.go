package main

import (
	"log"

	"github.com/phoihos/gosim/database"
	_ "github.com/phoihos/gosim/database/postgres"

	//_ "github.com/phoihos/gosim/database/mysql"
	//_ "github.com/phoihos/gosim/database/mssql"

	"github.com/phoihos/gosim/server"

	_ "handler"
)

func main() {
	defer database.Close()

	// Change orm naming rule if you want to change
	// database.SetOrmNamingStrategy(...)

	dbConf := &database.Configuration{Alias: "example", Host: "127.0.0.1", Port: "1433", Database: "exam", User: "user", Password: "password"}
	if err := database.OpenConnection(dbConf); err != nil {
		log.Print(err)
	}

	conf := &server.Configuration{Port: "8080", ShutdownPath: "/shutdown"}
	server.Run(conf)
}
