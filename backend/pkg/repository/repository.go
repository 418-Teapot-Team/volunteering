package repository

import (
	"fmt"
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
	ShareTask(taskId int, share bool, userId int) (err error)
	GetSharedTasks(userId int) (tasks []volunteering.TaskGetter, err error)

	// marks
	MarkAsDoneVolunteer(userId, taskId, trackedHours int) error
	MarkAsDoneEmployer(userId, taskId, trackedHours int, done bool) error

	// applies
	MakeApply(input *volunteering.Applies) (err error)
	GetAllApplies(userId int) (applies []volunteering.AppliesGetter, err error)
	ApproveApply(userId, Id, applyId int) error

	GetTimeStats(userId int) (data []volunteering.FinancialData, err error)
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

func (db *dbSQL) MarkAsDoneVolunteer(userId, taskId, tracked int) error {

	var trackedHours struct {
		Hours int `gorm:"column:estimate_time"`
	}

	if err := db.db.Table("tasks").Select("estimate_time").Where("assignee = ?", userId).Where("id = ?", taskId).First(&trackedHours).Error; err != nil {
		return err
	}

	return db.db.Model(&volunteering.Task{}).Where("assignee = ?", userId).Where("id = ?", taskId).
		Updates(map[string]interface{}{
			"pending":       true,
			"tracked_hours": tracked,
		}).Error
}

func (db *dbSQL) MarkAsDoneEmployer(userId, taskId, tracked int, done bool) error {
	var trackedHours struct {
		Assignee int `gorm:"column:assignee"`
		Hours    int `gorm:"column:estimate_time"`
	}

	if err := db.db.Table("tasks").Select("assignee, estimate_time").Where("user_id = ?", userId).Where("id = ?", taskId).First(&trackedHours).Error; err != nil {
		return err
	}

	fmt.Println(trackedHours)
	if done {
		user, _ := db.GetUserById(userId)
		if trackedHours.Assignee != userId && trackedHours.Assignee != 0 && user.Verified {
			err := db.db.Model(&volunteering.User{}).Where("id = ?", trackedHours.Assignee).Update("scores", tracked+trackedHours.Hours*2).Error
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		return db.db.Model(&volunteering.Task{}).Where("user_id = ? AND is_finished != true", userId).Where("id = ?", taskId).
			Updates(map[string]interface{}{
				"is_finished":   done,
				"pending":       false,
				"tracked_hours": tracked,
			}).Error
	}
	return db.db.Model(&volunteering.Task{}).Where("user_id = ?", userId).Where("id = ?", taskId).
		Updates(map[string]interface{}{
			"is_finished":   done,
			"pending":       true,
			"tracked_hours": 0,
		}).Error

}
