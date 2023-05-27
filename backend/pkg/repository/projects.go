package repository

import (
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

	err = tx.Where("id = ?", taskId).Where("user_id = ?", userId).Delete(&volunteering.Task{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
