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
		Preload("UserGetter").
		Find(&applies).Error
	if err != nil {
		return
	}

	return applies, nil
}
