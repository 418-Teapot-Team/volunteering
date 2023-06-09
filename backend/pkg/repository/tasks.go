package repository

import (
	"time"
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

func (db *dbSQL) GetTasks(userId int) ([]volunteering.TasksDB, error) {
	var tasks []volunteering.TasksDB

	projects, err := db.GetProjects(userId)
	if err != nil {
		return tasks, err
	}

	for _, project := range projects {
		var slice []volunteering.Task
		db.db.Where("project_id = ?", project.ID).Where("user_id = ?", userId).
			Preload("Project").
			Find(&slice)

		tasks = append(tasks, volunteering.TasksDB{Project: project, Task: slice})
	}
	return tasks, nil

}

func (db *dbSQL) GetUserTasks(userId int) (tasks []volunteering.Task, err error) {
	err = db.db.Order("created_at desc").Where("assignee = ?", userId).Where("is_finished != ?", true).Preload("User").Find(&tasks).Error
	if err != nil {
		return
	}

	return tasks, nil
}

func (db *dbSQL) ShareTask(taskId int, share bool, userId int) (err error) {
	var task volunteering.Task
	err = db.db.Where("id = ? AND user_id = ?", taskId, userId).First(&task).Error
	if err != nil {
		return err
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
		Select("DATE_FORMAT(closed_at, '%Y-%m-%d') AS date, IFNULL(SUM(tracked_hours),0) AS value").
		Where("user_id = ?", userId).
		Where("created_at >= ?", time.Now().AddDate(0, -1, 0)).
		Group("date").
		Order("STR_TO_DATE(date, \"%Y-%m-%d\")")

	err = query.Find(&data).Error

	if err != nil {
		return
	}

	return data, nil
}

func (db *dbSQL) GetGeneralStats(userId int) (result *volunteering.StatResult, err error) {

	type scoreStruct struct {
		scores int
	}

	scor := new(scoreStruct)

	type taskAmount struct {
		total int
	}

	tasks := new(taskAmount)

	type hoursAmount struct {
		total int
	}

	hours := new(hoursAmount)

	err = db.db.Where("id = ?", userId).Model(&volunteering.User{}).Select("IFNULL(scores, 0) as scores").Scan(&scor.scores).Error
	if err != nil {
		return
	}

	// total task amount for user
	err = db.db.Model(&volunteering.Task{}).Where("user_id = ?", userId).Select("count(*) as total").Find(&tasks.total).Error
	if err != nil {
		return
	}

	// total tracked hours for user
	err = db.db.Model(&volunteering.Task{}).Where("user_id = ?", userId).Select("IFNULL(sum(tracked_hours),0) as total").Find(&hours.total).Error
	if err != nil {
		return
	}

	return &volunteering.StatResult{Score: scor.scores, TaskAmount: tasks.total, HoursAmount: hours.total}, nil
}
