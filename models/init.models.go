package models

import (
	"anime-manager/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	conn, err := dbSetup()
	if err != nil {
		log.Fatalln(err)
	}
	doMigration(conn)
	return conn
}

func dbSetup() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.GetEnv("HOST"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_PORT"),
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return conn, err
}

func doMigration(db *gorm.DB) {
	db.AutoMigrate(Family{})
	db.AutoMigrate(Anime{})

	//if db != nil && db.Error != nil {
	//	log.Fatalln(db)
	//}
}
