package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"roomino/api"
	"roomino/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.Cors())
	v1 := r.Group("")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		v1.POST("register", api.UserRegisterHandler())
		//v1.POST("user/login", api.UserLoginHandler())
		//authed := v1.Group("/")
		/*authed.Use(middleware.JWT())
		{

			authed.POST("task_create", api.CreateTaskHandler())
			authed.GET("task_list", api.ListTaskHandler())
			authed.GET("task_show", api.ShowTaskHandler())
			authed.POST("task_update", api.UpdateTaskHandler())
			authed.POST("task_search", api.SearchTaskHandler())
			authed.POST("task_delete", api.DeleteTaskHandler())

		}*/
	}
	return r
}
