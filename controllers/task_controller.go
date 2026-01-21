package controllers

import (
	"net/http"
	"task-manager/models"
	"task-manager/repository"
	"task-manager/services"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val, _ := c.Get("user_id")
	userID := val.(uint)

	task.UserID = userID
	task.Status = "pending"

	if err := repository.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}

	services.TaskChannel <- task.ID
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	tasks, err := repository.GetAllTasks(userID.(uint), role.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := repository.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
