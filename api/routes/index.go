package routes

import (
	"net/http"

	ctr "github.com/eddyflawless/slack-service/api/controllers"
	mw "github.com/eddyflawless/slack-service/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())

	v1 := router.Group("/v1")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Healthy",
		})
	})

	router.Use(mw.Authenticate())

	g1 := v1.Group("/messages")

	{
		g1.POST("/", ctr.GetMessages)
		g1.POST("/process", ctr.ProcessMessages)
		g1.POST("/slack", ctr.SendSlackMessage)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Healthy",
		})
	})

	return router

}
