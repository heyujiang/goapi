package model

import "time"

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" `
	CreatedAt time.Time  `gorm:"column:created_at;default:null" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:null" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}
