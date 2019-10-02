package main

import "github.com/jinzhu/gorm"

func setupDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&MSG{})

	return db
}
