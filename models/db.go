package model

import (
	"fmt"

	"github.com/dev-parvej/go-blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(go-blog-db:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{}, &Post{})

	if err != nil {
		panic(err)
	}

	return db
}
