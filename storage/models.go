package storage

import "gorm.io/gorm"

// i know about relations ill consider it later
type GuildAccount struct {
	gorm.Model
	GuildID int64 `gorm:"guild_id"`
	UserID  int64 `gorm:"user_id"`
	Amount  int64 `gorm:"amount"`
}
