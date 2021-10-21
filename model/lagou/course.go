package lagou

import (
	"goapi/model"
	"goapi/pkg/client"
)

type CourseModel struct {
	model.BaseModel
	CourseId     int    `json:"courseId" gorm:"cloumn:course_id;not null" binding:"required"`
	Title        string `json:"title" gorm:"cloumn:title;not null" binding:"required"`
	Brief        string `json:"brief" gorm:"cloumn:brief;not null" binding:"required"`
	Image        string `json:"image" gorm:"cloumn:image;not null" binding:"required"`
	TeacherName  string `json:"teacherName" gorm:"cloumn:teacher_name;not null" binding:"required"`
	TeacherTitle string `json:"teacherTitle" gorm:"cloumn:teacher_title;not null" binding:"required"`
}

func (u *CourseModel) TableName() string {
	return "course"
}

func (c *CourseModel) Create() error {
	return client.MySqlClients.Self.Create(&c).Error
}
