package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	Self *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func (db *Database) Init()  {
	DB = &Database{
		Self:GetSelfDB(),
		Docker:GetDockerDB(),
	}
}

func (db *Database) Close() {
	_ = db.Self.Close()
	_ = db.Docker.Close()
}

func GetSelfDB() *gorm.DB  {
	return openDB(viper.GetString("database.host")+":"+viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.db"),)
}

func GetDockerDB() *gorm.DB  {
	return openDB(viper.GetString("docker_database.host")+":"+viper.GetString("docker_database.port"),
		viper.GetString("docker_database.user"),
		viper.GetString("docker_database.password"),
		viper.GetString("docker_database.db"),)
}

func openDB(addr,username,password,name string) *gorm.DB{
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	log.Println(config)

	db , err := gorm.Open("mysql",config)
	if err != nil {
		log.Fatal("Database connection failed. Database name: ",name," err : ", err)
	}

	setUp(db)

	return db
}

func setUp(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetConnMaxLifetime(100)
	db.DB().SetConnMaxIdleTime(1)
}

