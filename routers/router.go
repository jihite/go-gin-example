package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jihite/go-gin-example/docs"
	"github.com/jihite/go-gin-example/routers/api"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apig := r.Group("/api")
	{
		apig.GET("/tags", api.GetTags)
		apig.POST("/tags", api.AddTag)
		apig.PUT("/tags/:id", api.EditTag)
		apig.DELETE("/tags/:id", api.DeleteTag)
		apig.POST("/tags/export", api.ExportTag)
		apig.POST("/tags/import", api.ImportTag)
	}
	return r
}
