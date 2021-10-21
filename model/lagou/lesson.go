package lagou

import (
	"goapi/model"
	"goapi/pkg/client"
)

type LessonModel struct {
	model.BaseModel
	LessonId  int    `json:"lessonId" gorm:"cloumn:lesson_id;not null" binding:"required"`
	SectionId int    `json:"sectionId" gorm:"cloumn:section_id;not null" binding:"required"`
	CourseId  int    `json:"courseId" gorm:"cloumn:course_id;not null" binding:"required"`
	Theme     string `json:"theme" gorm:"cloumn:theme;not null" binding:"required"`
	Sort      int    `json:"sort" gorm:"cloumn:sort;not null" binding:"required"`
}

func (u *LessonModel) TableName() string {
	return "lesson"
}

func (l *LessonModel) Create() error {
	return client.MySqlClients.Self.Create(&l).Error
}
