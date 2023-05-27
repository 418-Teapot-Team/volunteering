package volunteering

import "time"

type Applies struct {
	ID            int64     `gorm:"column:id"`
	TaskID        int64     `gorm:"column:task_id"`
	RespondUserID int64     `gorm:"column:respond_user_id"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}
