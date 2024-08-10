package handler

import "github.com/gin-gonic/gin"

func GetHelloMessage() string {
	return "Hello from gin"
}

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": GetHelloMessage(),
	})
}

func PostHandler(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}

	err := c.ShouldBindJSON(&json)
	if err == nil {
		c.JSON(200, gin.H{"message": json.Message})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
