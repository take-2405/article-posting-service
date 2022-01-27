package middleware

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"prac-orm-transaction/config"
)

// Driver名
const DriverName = "mysql"

var DBConnectionInfo string

var Db *gorm.DB

func Init() error {
	var err error
	DBConnectionInfo = config.GetMysqlConnectionInfo()
	if err = createDBConnection(); err != nil {
		return err
	}
	log.Println("Successfull DB Connection")
	return err
}

func createDBConnection() error {
	var err error
	Db, err = gorm.Open(DriverName, DBConnectionInfo)
	if err != nil {
		return err
	}
	if err = Db.DB().Ping(); err != nil {
		return err
	}
	return nil
}
