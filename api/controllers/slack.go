package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSlackMessage(c *gin.Context) {
	var message string
	var err error
	if err = c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func SendTelegramMessage(c *gin.Context) {

}
