package main

import (
	"bwastartup/config"
	"bwastartup/routes"
	"log"
)

func main() {

	db, _ := config.ConnectDatabase()

	r := routes.SetupRouter(db)

	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// apply migration
	config.InitialMigration()


	err := r.Run("127.0.0.1:3000")
	if err != nil {
		log.Fatalln(err.Error())
	}

}
