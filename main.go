package main

import (
	"github.com/diego-arend/boilerplate_go_gin_gorm/database"
	"github.com/diego-arend/boilerplate_go_gin_gorm/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
