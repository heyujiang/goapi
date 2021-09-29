package model

import "time"

type BaseModel struct {
	Id       uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" `
	CreateAt time.Time  `gorm:"column:created_at;default:null" json:"-"`
	UpdateAt time.Time  `gorm:"column:updated_at;default:null" json:"-"`
	DeleteAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}
