package repository

import (
	"gorm.io/gorm"
	"time"
	"volunteering"
)

type Repository interface {
	CreateUser(user *volunteering.User) (err error)
	GetUserAuth(email, password string) (user volunteering.User, err error)
	UpdateLastLogin(userId int) error
}

type dbSQL struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) Repository {
	return &dbSQL{db: db}
}

func (db *dbSQL) CreateUser(user *volunteering.User) (err error) {
	tx := db.db.Begin()

	err = tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		if handleError(err) != nil {
			return err
		}
		return UniqueViolationError
	}
	return tx.Commit().Error
}

func (db *dbSQL) GetUserAuth(email, password string) (user volunteering.User, err error) {

	err = db.db.Where("email = ?", email).Where("password_hash = ?", password).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (db *dbSQL) UpdateLastLogin(userId int) error {
	return db.db.Model(&volunteering.User{}).Where("id = ?", userId).Update("last_login", time.Now()).Error
}
