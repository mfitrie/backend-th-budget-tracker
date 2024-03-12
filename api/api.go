package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router *gin.Engine) {
	route := router.Group("/api")
	route.GET("/", func(ctx *gin.Context) {
		// appContext := ctx.MustGet("_app").(gin.Context)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})

	})
}

// func AttachContext(app gin.Context) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("_app", app)
// 	}
// }
