package main

import (
	"stilyng94/fiber-crm/config"
	"stilyng94/fiber-crm/database"
)

func main() {
	config.SetupConfig()
	database.InitializeDatabase()
	defer database.Db.Close()
	database.Migrate()
	createServer()

}
