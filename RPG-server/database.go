package main

import "github.com/jinzhu/gorm"

func setupDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	db = db.Set("gorm:auto_preload", true)
	db.AutoMigrate(&DBUser{})
	db.AutoMigrate(&DBCharacter{})
	db.AutoMigrate(&DBItem{})
	db.AutoMigrate(&DBGame{})

	return db
}
