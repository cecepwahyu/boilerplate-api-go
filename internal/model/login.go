package model

import "time"

type Peserta struct {
	IDPeserta      uint      `gorm:"primaryKey;column:id_peserta"`
	Email          string    `gorm:"column:email;unique"`
	Password       string    `gorm:"column:password"`
	Token          string    `gorm:"column:token"`
	TokenUpdatedAt time.Time `gorm:"column:token_updated_at"`
	IsActive       bool      `gorm:"column:is_active"`
}

func (Peserta) TableName() string {
	return "tbl_peserta"
}
