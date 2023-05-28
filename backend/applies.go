package volunteering

import "time"

const appliesTable = "applies"

func (AppliesGetter) TableName() string {
	return appliesTable
}

type Applies struct {
	ID            int       `json:"-" gorm:"column:id"`
	TaskID        int       `json:"taskId" gorm:"column:task_id"`
	RespondUserID int       `json:"respondUserID" gorm:"column:respond_user_id"`
	AppliedUserID int       `json:"appliedUserID,omitempty" gorm:"column:applied_user_id"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
}

type AppliesGetter struct {
	Id         int        `json:"id" gorm:"column:id"`
	UserID     int        `json:"-" gorm:"column:applied_user_id"`
	TaskID     int        `json:"-" gorm:"column:task_id"`
	Task       Task       `json:"task" gorm:"foreignKey:TaskID"`
	UserGetter UserGetter `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"created_at"`
}
