package router

import (
	"ginseng_start/controller"
	"ginseng_start/middleware"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
/*func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", controllers.GetUsers)
		grp1.POST("user", controllers.CreateUser)
		grp1.GET("user/:id", controllers.GetUserByID)
		grp1.PUT("user/:id", controllers.UpdateUser)
		grp1.DELETE("user/:id", controllers.DeleteUser)
	}
	return r
}*/

// SetupRouter ... Configure routes
func SetupRouter(r *gin.Engine) {
	authMiddleware := middleware.NewJwt(r)

	r.GET("error", func(context *gin.Context) {
		panic("test")
	})

	r.POST("/login", authMiddleware.LoginHandler)

	// r.NoRoute(middleware.Jwt(authMiddleware), func(c *gin.Context) {
	// 	claims := jwt.ExtractClaims(c)
	// 	log.Printf("NoRoute claims: %#v\n", claims)
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	// r.Use(middleware.Jwt(authMiddleware))

	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", controller.GetUsers)
		grp1.POST("user", controller.CreateUser)
		grp1.GET("user/:id", controller.GetUserByID)
		grp1.PUT("user/:id", controller.UpdateUser)
		grp1.DELETE("user/:id", controller.DeleteUser)
	}

}
