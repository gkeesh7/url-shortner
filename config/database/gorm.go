package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var (
	once sync.Once
	db   *gorm.DB
	gerr error
)

var GetDbHandle = func(noPanic ...bool) *gorm.DB {
	once.Do(func() {
		db, gerr = gorm.Open("mysql", "root:@tcp(localhost:3306)/url_shortner?parseTime=true")
		if gerr != nil {
			log.Printf("The error while connecting to db %v", gerr.Error())
			panic("failed to connect database")
		}
		db.LogMode(true)
		//db.DB().ShowSQL(true)
		db.DB().SetMaxIdleConns(3)
		db.DB().SetMaxOpenConns(30)
		db.DB().SetConnMaxLifetime(-1)
	})
	return db
}
