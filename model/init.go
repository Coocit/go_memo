package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(conn string) {
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic("Mysql数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	db.SingularTable(true)       // 表名不加s，user
	db.DB().SetMaxIdleConns(20)  // 设置连接池
	db.DB().SetMaxOpenConns(100) // 最大连接数
	DB = db
	// migration()
}
