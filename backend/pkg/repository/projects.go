package repository

import (
	"log"
	"time"
	"volunteering"
)

func (db *dbSQL) CreateProject(input *volunteering.Project) (err error) {
	tx := db.db.Begin()

	err = tx.Create(input).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *dbSQL) DeleteProject(listId, userId int) (err error) {
	tx := db.db.Begin()

	err = tx.Where("id = ?", listId).Where("user_id = ?", userId).Delete(&volunteering.Project{}).Error
	if err != nil {
		return err
	}

	return tx.Commit().Error
}

func (db *dbSQL) GetProjects(userId int) (projects []volunteering.Project, err error) {
	err = db.db.Order("created_at desc").Where("user_id = ?", userId).Find(&projects).Error
	if err != nil {
		return
	}
	return projects, nil

}

func (db *dbSQL) DeleteTaskProject(userId int, taskId int) (err error) {
	tx := db.db.Begin()

	log.Println("DeleteTaskProject", taskId, userId)

	err = tx.Where("id = ? AND user_id = ?", taskId, userId).Delete(&volunteering.Task{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (db *dbSQL) GetProjectStats(userId int) (data []volunteering.ProjectData, err error) {
	err = db.db.Table("projects").
		Select("projects.title AS title, IFNULL(SUM(tasks.tracked_hours),0) AS value").
		Joins("LEFT JOIN tasks ON projects.id = tasks.project_id").
		Where("tasks.user_id = ?", userId).
		Where("projects.created_at >= ?", time.Now().AddDate(0, -1, 0)).
		Group("projects.title").
		Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}
