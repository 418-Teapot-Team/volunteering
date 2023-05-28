package repository

import "volunteering"

func (db *dbSQL) MakeApply(input *volunteering.Applies) (err error) {
	tx := db.db.Begin()

	err = tx.Create(input).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *dbSQL) GetAllApplies(userId int) (applies []volunteering.AppliesGetter, err error) {
	err = db.db.Order("created_at desc").Where("applied_user_id = ?", userId).
		Preload("Task").
		Preload("AppliedUserGetter").
		Preload("Task.Project").
		Find(&applies).Error
	if err != nil {
		return
	}

	return applies, nil
}

func (db *dbSQL) ApproveApply(userId, Id, applyId int) (err error) {
	tx := db.db.Begin()
	err = tx.Model(volunteering.Applies{}).Where("respond_user_id = ? AND applied_user_id = ?", userId, applyId).Where("id = ?", Id).Update("accepted", true).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var result struct {
		Id int `json:"id" gorm:"task_id"`
	}

	err = tx.Model(volunteering.Applies{}).Where("respond_user_id = ? AND applied_user_id = ?", userId, applyId).Where("id = ?", Id).Select("task_id as id").Find(&result).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&volunteering.Task{}).Where("user_id = ?", userId).Where("id = ?", result.Id).
		Updates(map[string]interface{}{
			"is_finished": false,
			"pending":     true,
			"assignee":    applyId,
		}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
