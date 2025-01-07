package storage

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func New() *Storage {
	dsn := os.Getenv("MYSQL_DSN")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("failed to init storage:", err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&GuildAccount{})

	return &Storage{
		DB: db,
	}
}

func (s *Storage) UserExists(id int64) bool {
	return s.DB.First(&User{
		ID: id,
	}).RowsAffected > 0
}
