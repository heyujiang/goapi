package lagou

import (
	"goapi/model"
	"goapi/pkg/client"
)

type ContentModel struct {
	model.BaseModel
	LessonId int    `json:"lessonId" gorm:"cloumn:lesson_id;not null" binding:"required"`
	Content  string `json:"content" gorm:"cloumn:content" binding:"required"`
}

func (u *ContentModel) TableName() string {
	return "content"
}

func (c *ContentModel) Create() error {
	return client.MySqlClients.Self.Create(&c).Error
}
