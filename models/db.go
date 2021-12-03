package models

import (
	"blog/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var db *gorm.DB

func Init() {
	db, _ = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	// 连接池中最大的闲置连接数
	db.DB().SetMaxIdleConns(10)
	// 最大连接数
	db.DB().SetMaxOpenConns(100)
	//最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)
}
