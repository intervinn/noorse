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

	db.AutoMigrate(&GuildAccount{})

	return &Storage{
		DB: db,
	}
}

func (s *Storage) UserExists(guild int64, user int64) bool {
	rows := s.DB.Where("guild_id = ? AND user_id = ?", guild, user).RowsAffected
	fmt.Println(rows)
	return rows > 0
}
