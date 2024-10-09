package main

import (
	"awesomeProject/cmd/api"
	"awesomeProject/config"
	"awesomeProject/db"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	database, err := db.NewSqlDB(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddres,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(database)
	server := api.NewApiServer(":8080", database)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to DATABASE")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
