package repository

import (
	"volunteering"
)

func (db *dbSQL) CreateTask(input *volunteering.Task) (err error) {
	tx := db.db.Begin()
	err = tx.Create(input).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *dbSQL) DeleteTask(taskId, userId int) (err error) {
	tx := db.db.Begin()

	err = tx.Model(&volunteering.Task{}).Where("id = ?", taskId).Where("assignee = ?", userId).Update("assignee", nil).Update("shared", true).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (db *dbSQL) GetTasks(userId int) (tasks []volunteering.TasksDB, err error) {

	err = db.db.Order("created_at desc").Where("user_id = ?", userId).
		Preload("Project").
		Preload("Task").
		Preload("Task.Project").
		Find(&tasks).Error
	if err != nil {
		return
	}

	return tasks, nil

}

func (db *dbSQL) GetUserTasks(userId int) (tasks []volunteering.Task, err error) {
	err = db.db.Order("created_at desc").Where("assignee = ?", userId).Find(&tasks).Error
	if err != nil {
		return
	}

	return tasks, nil
}

func (db *dbSQL) ShareTask(taskId int, share bool, userId int) (err error) {
	var task volunteering.Task
	err = db.db.Where("id = ? AND user_id = ?", taskId, userId).First(&task).Error
	if err != nil {
		return
	}

	task.Shared = share

	err = db.db.Save(&task).Error
	if err != nil {
		return
	}

	return nil
}

func (db *dbSQL) GetSharedTasks(userId int) (tasks []volunteering.TaskGetter, err error) {
	err = db.db.Order("created_at desc").Where("assignee is NULL AND shared=true AND is_finished=false").Preload("User").Preload("Project").Find(&tasks).Error
	if err != nil {
		return
	}

	return tasks, nil
}

func (db *dbSQL) GetTimeStats(userId int) (data []volunteering.FinancialData, err error) {
	query := db.db.Table("tasks").
		Select("DATE_FORMAT(closed_at, '%Y-%m') AS date, SUM tracked_hours AS value").
		Where("user_id = ?", userId).
		Group("date").
		Order("STR_TO_DATE(date, \"%Y-%m\")")

	err = query.Find(&data).Error

	if err != nil {
		return
	}

	return data, nil
}
