package initial

import (
	"github.com/artemxgod/gorm-tutorial/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB(databaseURL string) (*gorm.DB, error) {
	// initialize database using database url
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// get mysql database to ping 
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}

	// ping db to check connection
	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	// migrate models into database
	if err := db.AutoMigrate(&model.User{}, &model.BankAccount{}); err != nil {
		return nil, err
	}

	return db, nil
}