package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/xingxingso/ginseng_start/model"
	"github.com/xingxingso/ginseng_start/service"

	"github.com/sirupsen/logrus"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"
var userNameKey = "name"

func NewJwt() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "gin jwt",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			logrus.Infof("PayloadFunc %T, %+[1]v", data)
			if v, ok := data.(*model.User); v != nil && ok {
				return jwt.MapClaims{
					identityKey: v.Id,
					userNameKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			logrus.Infof("IdentityHandler %T,%+[1]v", claims)
			return &model.User{
				Id:   uint(claims[identityKey].(float64)),
				Name: claims[userNameKey].(string),
			}
		},
		// 登录
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login service.Credential
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := login.Login()
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		// 登录响应格式
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Unix(),
			})
		},
		// 刷新响应
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Unix(),
			})
		},

		// 鉴权
		Authorizator: func(data interface{}, c *gin.Context) bool {
			logrus.Infof("Authorizator %T, %+[1]v", data)
			// if v, ok := data.(*model.User); ok && v.Name == "admin" {
			if v, ok := data.(*model.User); ok && v.Id == 1 {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			logrus.Error("Unauthorized")
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	// errInit := authMiddleware.MiddlewareInit()

	// if errInit != nil {
	// 	log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	// 	return nil
	// }

	return authMiddleware
}

func Jwt(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return authMiddleware.MiddlewareFunc()
}
