package repository

import (
	"task-manager/config"
	"task-manager/models"
)

func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	err := config.DB.First(&task, "id = ?", id).Error
	return task, err
}

func GetAllTasks(userID uint, role string) ([]models.Task, error) {
	var tasks []models.Task
	if role == "admin" {
		// Admin sees all [cite: 35]
		err := config.DB.Find(&tasks).Error
		return tasks, err
	}
	// User sees only their own [cite: 34]
	err := config.DB.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func DeleteTask(id string) error {
	return config.DB.Delete(&models.Task{}, "id = ?", id).Error
}

func UpdateTaskStatus(id string, status string) error {
	return config.DB.Model(&models.Task{}).Where("id = ?", id).Update("status", status).Error
}
