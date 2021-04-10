package mysql

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	//TODO
	db, err = gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/forum?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err : ", err)
		os.Exit(0)
	} else {
		fmt.Println("mysql forum database Open success,", db)
	}
	db.Debug()
}

func GetDB() *gorm.DB {
	return db
}
