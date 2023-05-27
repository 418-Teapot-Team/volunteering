package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(connection string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
