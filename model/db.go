package model

import (
	"doovvvblog/utils"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDSN() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.AppConfig.Database.Username,
		utils.AppConfig.Database.Password,
		utils.AppConfig.Database.Host,
		utils.AppConfig.Database.Port,
		utils.AppConfig.Database.Dbname)
}

var DB *gorm.DB
var err error

func InitDB() {
	DB, err = gorm.Open(mysql.Open(getDSN()), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(10 * time.Second)

	DB.AutoMigrate(&User{}, &Category{}, &Article{})
}
