package router

import (
	"log"
	"net/http"

	"github.com/xingxingso/ginseng_start/controller"
	"github.com/xingxingso/ginseng_start/middleware"
	"github.com/xingxingso/ginseng_start/model"

	"github.com/sirupsen/logrus"

	jwt "github.com/appleboy/gin-jwt/v2"

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
	authMiddleware := middleware.NewJwt()

	r.GET("error", func(context *gin.Context) {
		panic("test")
	})

	r.POST("/login", authMiddleware.LoginHandler)
	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	r.NoRoute(func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	jwtMid := middleware.Jwt(authMiddleware)

	grp1 := r.Group("/user-api")
	grp1.Use(jwtMid)
	{
		grp1.GET("user", controller.GetUsers)
		grp1.POST("user", controller.CreateUser)
		grp1.GET("user/:id", controller.GetUserByID)
		grp1.PUT("user/:id", controller.UpdateUser)
		grp1.DELETE("user/:id", controller.DeleteUser)
	}

	auth := r.Group("/auth")
	// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(jwtMid)
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get(authMiddleware.IdentityKey)
			logrus.Infof("%+v", user)
			c.JSON(http.StatusOK, gin.H{
				"userID":   claims[authMiddleware.IdentityKey],
				"userName": user.(*model.User).Name,
				"text":     "Hello World.",
			})
		})
	}
}
