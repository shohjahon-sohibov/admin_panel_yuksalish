package api

import (
	"freelance/admin_panel/api/handlers"
	"freelance/admin_panel/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter godoc
// @description This is a api gateway
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(customCORSMiddleware())

	// BRANCH
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch", h.GetBranchList)
	r.GET("/branch/:branch_id", h.GetSingleBranch)
 	r.PUT("/branch/:branch_id", h.UpdateBranch)
 	r.DELETE("/branch/:branch_id", h.DeleteBranch)
 
	// GROUP
	r.POST("/group", h.CreateGroup)
	r.GET("/group", h.GetGroupList)
	r.GET("/group/:group_id", h.GetSingleGroup)
	r.PUT("/group/:group_id", h.UpdateGroup)
	r.DELETE("/group/:group_id", h.DeleteGroup)

	// STUDENT
	r.POST("/student", h.CreateStudent)
	r.GET("/student", h.GetStudentList)
	r.GET("/student/:student_id", h.GetSingleStudent)
	r.PUT("/student/:student_id", h.UpdateStudent)
	r.DELETE("/student/:student_id", h.DeleteStudent)


	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

		return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}