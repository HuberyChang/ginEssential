package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	pasword := "root"
	charset := "utf8mb4"
	loc := "Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		pasword,
		host,
		port,
		database,
		charset,
		loc)
	//viper.Get("dataSource.user"),
	//viper.Get("dataSource.password"),
	//viper.Get("dataSource.host"),
	//viper.Get("dataSource.port"),
	//viper.Get("dataSource.database"),
	//viper.Get("dataSource.charset"),
	//viper.Get("dataSource.loc"),

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	DB.Debug()
	return DB
}
