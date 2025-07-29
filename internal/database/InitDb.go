package database

import (
	"fmt"

	"github.com/rest_api_motivation/internal/config"
	"github.com/rest_api_motivation/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func Init(cfg config.Config) *DB {
	path := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.Db.User, cfg.Db.Password, cfg.Db.NameDB, cfg.Db.PortDB)
	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Post{})

	return &DB{
		db: db,
	}
}
