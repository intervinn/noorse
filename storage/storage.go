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

func Paginate(per, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page) * per
		return db.Offset(offset).Limit(per)
	}
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

func (s *Storage) Paginate(page int) *gorm.DB {
	return s.DB.Scopes(Paginate(10, page))
}

func (s *Storage) UserExists(guild int64, user int64) bool {
	var c int64
	s.DB.
		Model(&GuildAccount{}).
		Where("user_id = ? AND guild_id = ?", user, guild).
		Count(&c)
	return c != 0
}
