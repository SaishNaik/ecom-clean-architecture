package datastore

import (
	"ecom_clean_architecture/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func NewDB() *gorm.DB {
	DBMS := "mysql"
	mysqlConfig := &mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(DBMS, mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	return db

}
