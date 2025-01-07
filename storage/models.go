package storage

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64          `gorm:"primaryKey"`
	Accounts []GuildAccount `gorm:"many2many:user_accounts"`
}

type GuildAccount struct {
	gorm.Model
	ID     int64 `gorm:"primaryKey"`
	Amount int64
}
