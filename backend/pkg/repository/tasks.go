package repository

import (
	"log"
	"volunteering"
)

func (db *dbSQL) CreateTask(input *volunteering.Task) (err error) {
	tx := db.db.Begin()

	log.Println(input.ProjectId)

	err = tx.Create(input).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *dbSQL) DeleteTask(taskId, userId int) (err error) {
	tx := db.db.Begin()

	err = tx.Where("id = ?", taskId).Where("user_id = ?", userId).Delete(&volunteering.Task{}).Error
	if err != nil {
		return err
	}

	return tx.Commit().Error
}
