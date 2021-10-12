package model

import (
	"github.com/spf13/viper"
	"goapi/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"cloumn:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"cloumn:password;not null" binding:"required" validate:"min=5,max=128"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	SayHello  string `json:"say_hello"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

//创建新用户
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

//删除用户
func (u *UserModel) Delete() error {
	return DB.Self.Delete(&u).Error
}

//更新用户
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

//根据主键id获得用户信息
func GetUser(id uint64) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

//根据用户名称获得用户信息
func GetUserByUserName(username string) (model *UserModel, err error) {
	model = &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&model)
	err = d.Error
	return
}

//所有用户
func ListUser(offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = viper.GetInt("db_default_limit")
	}

	user := make([]*UserModel, 0)
	var count uint64

	if err := DB.Self.Model(&UserModel{}).Count(&count).Error; err != nil { //查询总条数
		return user, count, err
	}

	if err := DB.Self.Offset(offset).Limit(limit).Order("id desc").Find(&user).Error; err != nil {
		return user, count, err
	}

	return user, count, nil
}

//验证参数
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

//加密密码
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

//比较密码
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}
