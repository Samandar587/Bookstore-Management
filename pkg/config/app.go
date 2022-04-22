package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:Sawka@7011/dunyo?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic("Failed to connect to database.")
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
