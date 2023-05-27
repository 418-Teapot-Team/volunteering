package repository

import (
	"gorm.io/gorm"
	"time"
	"volunteering"
)

type Repository interface {
	// Auth
	CreateUser(user *volunteering.User) (err error)
	GetUserAuth(email, password string) (user volunteering.User, err error)
	GetUserById(userId int) (user volunteering.User, err error)

	UpdateLastLogin(userId int) error

	// Projects
	CreateProject(input *volunteering.Project) (err error)
	DeleteProject(listId, userId int) (err error)
	GetProjects(userId int) (projects []volunteering.Project, err error)
	DeleteTaskProject(userId int, taskId int) (err error)

	// Tasks
	CreateTask(input *volunteering.Task) (err error)
	DeleteTask(taskId, userId int) (err error)
	GetTasks(userId int) (tasks []volunteering.TasksDB, err error)
	GetUserTasks(userId int) (tasks []volunteering.Task, err error)

	MarkAsDoneVolunteer(userId, taskId int) error
	MarkAsDoneEmployer(userId, taskId int) error
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

func (db *dbSQL) GetUserById(userId int) (user volunteering.User, err error) {
	err = db.db.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *dbSQL) MarkAsDoneVolunteer(userId, taskId int) error {

	var trackedHours struct {
		TrackedHours int `gorm:"column:tracked_hours"`
		Hours        int `gorm:"column:hours"`
	}

	if err := db.db.Table("tasks").Select("tracked_hours, hours").Where("assignee = ?", userId).Where("id = ?", taskId).First(&trackedHours).Error; err != nil {
		return err
	}

	return db.db.Model(&volunteering.Task{}).Where("assignee = ?", userId).Where("id = ?", taskId).
		Updates(map[string]interface{}{
			"pending":       true,
			"tracked_hours": trackedHours.Hours + trackedHours.TrackedHours,
		}).Error
}

func (db *dbSQL) MarkAsDoneEmployer(userId, taskId int) error {
	return db.db.Model(&volunteering.Task{}).Where("user_id = ?", userId).Where("id = ?", taskId).
		Updates(map[string]interface{}{
			"is_finished": true,
		}).Error
}
