package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:1234567890@/forum?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err : ", err)
		os.Exit(0)
	} else {
		fmt.Println("mysql forum database Open success,", db)
	}
	db.Debug()
}