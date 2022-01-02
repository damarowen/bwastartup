
package config

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)



func ConnectDatabase() (db *gorm.DB, err error){
	dbUser := os.Getenv("DB_USER")

	fmt.Println(dbUser, "xxxxxxxx")

	dsn := "root:root@tcp(127.0.0.1:3306)/bwagolang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err  = gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	log.Println("CONNECTED TO DB")

	return db, err
}

func InitialMigration() {

	db, _ := ConnectDatabase()

	err := db.AutoMigrate(&campaign.Campaign{}, &user.User{}, &campaign.CampaignImage{})

	if err != nil {
		log.Println(err.Error())
		panic("failed to migrate database")
	}


}

