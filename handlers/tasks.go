package handlers

import (
	"fmt"
	"net/http"
	"package/db"

	"github.com/gin-gonic/gin"
)

// хендлер создания таски
func (h BaseHandler) CreateTask(c *gin.Context) {
	var task *db.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	userUUID, err := getUserUUIDFromToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	task.User = &db.User{UserId: userUUID}

	task, err = h.db.CreateTask(task)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// хендлер удаления таски
func (h BaseHandler) DeleteTask(c *gin.Context) {
	var task *db.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	userUUID, err := getUserUUIDFromToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	task.User = &db.User{UserId: userUUID}

	_, err = h.db.DeleteTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// хендлер получения всех тасок
func (h BaseHandler) GetAllTasks(c *gin.Context) {
	cards, err := h.db.GetAllTasks()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, cards)
}
