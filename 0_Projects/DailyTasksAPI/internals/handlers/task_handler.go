package handlers

import (
	"github/laxman67/internals/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func WelcomeRouteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "Home Route /GET",
		Success: true,
		Data:    nil,
	})
}

func GetTasksHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks []models.Task

		if err := db.Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Message: "Failed to fetch tasks",
				Success: false,
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Message: "All tasks fetched successfully",
			Success: true,
			Data:    tasks,
		})
	}
}

func GetTaskByIDHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		var task models.Task

		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, Response{
				Message: "Task not found",
				Success: false,
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Message: "Task fetched successfully",
			Success: true,
			Data:    task,
		})

	}
}

func CreateTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTask models.Task

		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Message: "Invalid JSON body",
				Success: false,
				Data:    nil,
			})
			return
		}

		if err := db.Create(&newTask).Error; err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Message: "Failed to create task",
				Success: false,
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, Response{
			Message: "Task added successfully!",
			Success: true,
			Data:    newTask,
		})
	}
}

func DeleteTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		result := db.Delete(&models.Task{}, id)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Message: "Failed to delete task",
				Success: false,
				Data:    nil,
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, Response{
				Message: "Task not found",
				Success: false,
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Message: "Task deleted successfully",
			Success: true,
			Data:    nil,
		})
	}
}
