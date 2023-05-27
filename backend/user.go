package volunteering

import "time"

type User struct {
	Id        int       `json:"-" gorm:"column:id"`
	FirstName string    `json:"firstName" binding:"required" gorm:"column:first_name"`
	LastName  string    `json:"lastName" binding:"required" gorm:"column:last_name"`
	Email     string    `json:"email" binding:"required" gorm:"column:email"`
	Password  string    `json:"password" binding:"required" gorm:"column:password_hash"`
	CreatedAt time.Time `json:"-" gorm:"created_at"`
	LastLogin time.Time `json:"lastLogin,omitempty" gorm:"last_login"`
}
