package routes

import (
	"github.com/Adejare77/blogPlatform/internals/controllers"
	"github.com/gin-gonic/gin"
)

var UserRoutes = func(r *gin.RouterGroup) {
	r.POST("/user/register", controllers.Register)
	r.POST("/user/login", controllers.Login)
	r.GET("/index", controllers.GetAllPosts)
	r.GET("/posts/:post_id", controllers.GetPostByID)
}
