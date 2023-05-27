package volunteering

import "time"

const (
	tasksTable = "tasks"
)

func (Task) TableName() string {
	return tasksTable

}

type Task struct {
	ID           int       `json:"id,omitempty" gorm:"column:id"`
	UserId       int       `json:"-" gorm:"column:user_id"`
	Assignee     *int      `json:"-" gorm:"column:assignee"`
	ProjectId    int       `json:"projectId" gorm:"column:project_id"`
	Title        string    `json:"title" gorm:"column:title" binding:"required"`
	Shared       bool      `json:"shared" gorm:"column:shared"`
	EstimateTime int       `json:"estimate_time,omitempty" gorm:"column:estimate_time"`
	IsFinished   bool      `json:"is_finished" gorm:"column:is_finished"`
	Description  string    `json:"description" gorm:"column:description"`
	CreatedAt    time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
}

type TasksDB struct {
	Project Project `json:"project" gorm:"foreignKey:ProjectID"`
	Tasks   []Task  `json:"tasks" gorm:"foreignKey:ProjectID"`
}
