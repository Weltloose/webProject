package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	// don't use := cause it won't change the global variable Db. HUGE HOLE
	Db, err = gorm.Open("mysql", "weltloose:Linux2.0@/homework_system?charset=utf8&parseTime=True&loc=Local")
	// And don't close here
	if err != nil {
		fmt.Println("data not openable: ", err)
		return
	}
	Db.AutoMigrate(&Ta{})
	Db.AutoMigrate(&Homework{})
}
