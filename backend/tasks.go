package volunteering

import "time"

const (
	tasksTable = "tasks"
)

func (Task) TableName() string {
	return tasksTable
}

func (TasksDB) TableName() string {
	return tasksTable
}

func (TaskGetter) TableName() string {
	return tasksTable
}

type Task struct {
	ID           int       `json:"id,omitempty" gorm:"column:id"`
	UserId       int       `json:"-" gorm:"column:user_id"`
	Assignee     *int      `json:"-" gorm:"column:assignee"`
	ProjectId    int       `json:"projectId,omitempty" gorm:"column:project_id"`
	Project      Project   `json:"project,omitempty" gorm:"foreignKey:ProjectId"`
	User         User      `json:"user,omitempty" gorm:"foreignKey:UserId"`
	Title        string    `json:"title" gorm:"column:title"`
	Shared       bool      `json:"shared" gorm:"column:shared"`
	EstimateTime int       `json:"estimate_time,omitempty" gorm:"column:estimate_time"`
	IsFinished   bool      `json:"is_finished" gorm:"column:is_finished"`
	Description  string    `json:"description,omitempty" gorm:"column:description"`
	CreatedAt    time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
}

type TaskGetter struct {
	Id           int        `json:"id,omitempty" gorm:"column:id"`
	UserId       int        `json:"-" gorm:"column:user_id"`
	User         UserGetter `json:"user" gorm:"foreignKey:UserId"`
	ProjectId    int        `json:"projectId" gorm:"column:project_id"`
	Project      Project    `json:"project" gorm:"foreignKey:ProjectId"`
	Title        string     `json:"title" gorm:"column:title"`
	EstimateTime int        `json:"estimate_time,omitempty" gorm:"column:estimate_time"`
	Description  string     `json:"description,omitempty" gorm:"column:description"`
	CreatedAt    time.Time  `json:"createdAt,omitempty" gorm:"column:created_at"`
}

//type TasksDB struct {
//	ProjectID int            `json:"-" gorm:"column:project_id"`
//	Project   Project        `json:"project" gorm:"foreignKey:ProjectID"`
//	Tasks     map[int][]Task `json:"tasks" gorm:"-"`
//}

type TasksDB struct {
	ProjectID int     `json:"-" gorm:"column:project_id"`
	ID        int     `json:"-" gorm:"column:id"`
	Project   Project `json:"project" gorm:"foreignKey:ProjectID"`
	Task      []Task  `json:"tasks" gorm:"foreignKey:ID"`
}

type FinancialData struct {
	Date  string
	Value float64
}

type StatResult struct {
	Score       int
	TaskAmount  int
	HoursAmount int
}
