package handler

import (
	"net/http"
	"strconv"
	"time"
	"training-golang/session-2-latihan-crud-user-gin/entity"

	"github.com/gin-gonic/gin"
)

var (
	users  []entity.User
	nextID int
)

func CreatedUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = nextID
	nextID++
	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()

	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, u := range users {
		if u.ID == id {
			updateUser := entity.User{
				ID:        id,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				CreatedAt: u.CreatedAt,
				UpdateAt:  time.Now(),
			}

			users[i] = updateUser
			c.JSON(http.StatusOK, updateUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
