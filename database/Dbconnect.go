package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"booking_system/utils"
)

var DbConnect *gorm.DB
var err error

func Db() {
	config, _ := utils.GetConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config["POSTGRES_DB_HOST"],
		config["POSTGRES_DB_PORT"],
		config["POSTGRES_DB_USER"],
		config["POSTGRES_DB_NAME"],
		config["POSTGRES_DB_PASSWD"],
	)

	// dsn := "host=192.168.56.102 user=lilee password=Lilee1234 dbname=postgres port=5432 sslmode=disable"
	DbConnect, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

}
