package volunteering

import "time"

const (
	projectsTable = "projects"
)

func (Project) TableName() string {
	return projectsTable

}

type Project struct {
	ID          int       `json:"id,omitempty" gorm:"column:id"`
	UserId      int       `json:"-" gorm:"column:user_id"`
	Title       string    `json:"title" gorm:"column:title" binding:"required"`
	Description string    `json:"description" gorm:"column:description"`
	CreatedAt   time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
}
