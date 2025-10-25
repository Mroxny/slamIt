package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/Mroxny/slamIt/internal/config"
	"github.com/Mroxny/slamIt/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var cfg = config.GetConfig().DB
var localDBPath = cfg.SQLitePath

func InitDB(localOnly bool) *gorm.DB {
	var db *gorm.DB
	var err error
	if localOnly {
		db, err = gorm.Open(sqlite.Open(localDBPath), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to database:", err)
		}
	} else {
		db, err = gorm.Open(sqlite.Open(localDBPath), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to database:", err)
		}
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Slam{},
		&model.Stage{},
		&model.Participation{},
		&model.Performance{},
		&model.Vote{},
	); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	return db
}

func ClearLocalDB() {
	err := os.Remove(localDBPath)
	if err != nil {
		fmt.Println("Error deleting local database: ", err)
		return
	}

	_, err = gorm.Open(sqlite.Open(localDBPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to a local database:", err)
	}
}
