package main

import (
	"example.com/project/cmd"
	"example.com/project/pkg/config"
	"example.com/project/pkg/database"
	"example.com/project/pkg/logger"
	"example.com/project/pkg/server"
)

func main() {
	config.InitConfig()
	database.InitDatabase()
	logger.InitLogger()
	server.StartServer()
}