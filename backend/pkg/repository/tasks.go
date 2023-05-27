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
	projects, err := db.GetProjects(userId)
	if err != nil {
		return tasks, err
	}

	var task []volunteering.Task

	for _, project := range projects {
		err = db.db.Order("created_at desc").Where("user_id = ?", userId).
			Where("project_id = ?", project.ID).
			Find(&task).Error
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, volunteering.TasksDB{Project: project, Tasks: task})
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
