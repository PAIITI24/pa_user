package model

import "time"

type Token struct {
	Token     string    `packets:"token" gorm:"PrimaryKey;not null"`
	ExpiredAt time.Time `packets:"expired_at"`
	Owner     int       `packets:"owner"`
	IsEnabled bool      `packets:"is_enabled"`
	CreatedAt time.Time `packets:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `packets:"updated_at" gorm:"autoUpdateTime"`
	User      User      `gorm:"foreignKey:Owner;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
