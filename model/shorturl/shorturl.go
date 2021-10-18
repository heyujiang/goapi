package shorturl

import (
	"goapi/model"
	"goapi/pkg/client"
)

type ShorturlModel struct {
	model.BaseModel
	LongUrl  string `json:"long_url" gorm:"cloumn:long_url;not null" binding:"required" validate:"min=1,max=32"`
	ShortUrl string `json:"short_url" gorm:"cloumn:short_url;not null" binding:"required" validate:"min=5,max=128"`
}

func (s *ShorturlModel) TableName() string {
	return "tb_shorturl"
}

func (s *ShorturlModel) Create() error {
	return client.MySqlClients.Self.Create(&s).Error
}

func GetInfoByShortUrl(shortUrl string) (*ShorturlModel, error) {
	s := &ShorturlModel{}
	d := client.MySqlClients.Self.Where("short_url = ?", shortUrl).First(&s)
	return s, d.Error
}
