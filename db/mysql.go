package db

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"theses/topic"
)

var db *gorm.DB

func connect() (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_POST"),
		os.Getenv("DB_NAME"),
	)

	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return errors.New("failed to connect database " + err.Error())
	}

	// Migrate the schema
	if err = db.AutoMigrate(&topic.Topic{}); err != nil {
		return errors.New("failed to migrate database " + err.Error())
	}

	return nil
}

func Con() *gorm.DB {
	if db == nil {
		if err := connect(); err != nil {
			panic(err)
		}
	}
	return db
}
