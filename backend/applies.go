package volunteering

import "time"

const appliesTable = "applies"

func (AppliesGetter) TableName() string {
	return appliesTable
}

func (AppliedUserGetter) TableName() string {
	return "users"
}

type Applies struct {
	ID            int       `json:"-" gorm:"column:id"`
	TaskID        int       `json:"taskId" gorm:"column:task_id"`
	RespondUserID int       `json:"respondUserID" gorm:"column:respond_user_id"`
	AppliedUserID int       `json:"appliedUserID,omitempty" gorm:"column:applied_user_id"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
}

type AppliedUserGetter struct {
	Id        int    `json:"id" gorm:"column:id"`
	FirstName string `json:"firstName" binding:"required" gorm:"column:first_name"`
	LastName  string `json:"lastName" binding:"required" gorm:"column:last_name"`
	Email     string `json:"email" binding:"required" gorm:"column:email"`
}

type AppliesGetter struct {
	Id         int        `json:"id" gorm:"column:id"`
	UserID     int        `json:"-" gorm:"column:applied_user_id"`
	TaskID     int        `json:"-" gorm:"column:task_id"`
	Task       Task       `json:"task" gorm:"foreignKey:TaskID"`
	UserGetter UserGetter `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"created_at"`
}
