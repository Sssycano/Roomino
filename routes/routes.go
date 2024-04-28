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

	// 定义顶级路由组
	v1 := r.Group("")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "success"})
		})

		v1.POST("register", api.UserRegisterHandler())
		v1.POST("login", api.UserLoginHandler())

		// 使用 JWT 中间件保护子组
		authed := v1.Group("", middleware.JWT()) // 确保 JWT 中间件应用到子组
		{
			authed.POST("profile/unitinfo", api.UnitInfoHandler()) // 在受保护的组中添加路由
		}
	}

	return r
}
