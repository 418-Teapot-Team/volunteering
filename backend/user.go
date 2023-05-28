package volunteering

import "time"

const (
	usersTable = "users"
)

func (User) TableName() string {
	return usersTable
}

func (UserGetter) TableName() string {
	return usersTable
}

type User struct {
	Id        int       `json:"-" gorm:"column:id"`
	Verified  bool      `json:"verified,omitempty" gorm:"column:verified"`
	FirstName string    `json:"firstName" gorm:"column:first_name"`
	LastName  string    `json:"lastName" gorm:"column:last_name"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" binding:"required" gorm:"column:password_hash"`
	CreatedAt time.Time `json:"-" gorm:"created_at"`
	LastLogin time.Time `json:"lastLogin,omitempty" gorm:"last_login"`
}

type UserGetter struct {
	Id        int    `json:"userId" gorm:"column:id"`
	Verified  bool   `json:"verified" gorm:"column:verified"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Email     string `json:"email" gorm:"column:email"`
}
