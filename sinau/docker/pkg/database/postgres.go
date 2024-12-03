package database

import (
	"ivandhitya/docker/internal/domain/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(host string, username string, password string, dbname string, port string) (*gorm.DB, error) {
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&entity.Student{}); err != nil {
		return nil, err
	}

	log.Println("Database connected and migrated")
	return db, nil
}
