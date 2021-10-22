package lagou

import (
	"goapi/model"
	"goapi/pkg/client"
)

type SectionModel struct {
	model.BaseModel
	SectionId   int    `json:"sectionId" gorm:"cloumn:section_id;not null" binding:"required"`
	CourseId    int    `json:"courseId" gorm:"cloumn:course_id;not null" binding:"required"`
	SectionName string `json:"sectionName" gorm:"cloumn:section_name;not null" binding:"required"`
	Sort        int    `json:"sort" gorm:"cloumn:sort;not null" binding:"required"`
	Description string `json:"description" gorm:"cloumn:description;not null" binding:"required"`
}

func (u *SectionModel) TableName() string {
	return "section"
}

func (s *SectionModel) Create() error {
	return client.MySqlClients.Self.Create(&s).Error
}

func ListSection(courseId int) ([]SectionModel, error) {
	var s []SectionModel
	if err := client.MySqlClients.Self.Where("course_id = ?", courseId).Find(&s).Error; err != nil {
		return nil, err
	}

	return s, nil
}

func GetSection(id uint64) (*SectionModel, error) {
	u := &SectionModel{}
	d := client.MySqlClients.Self.Where("section_id = ?", id).First(&u)
	return u, d.Error
}
