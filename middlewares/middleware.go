package middlewares

import (
	"context"
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"hong-blog/db"
	"hong-blog/httpService"
	"hong-blog/utils"
	"log"
	"net/http"
	"time"
)

var identityKey = "id"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	Id       int
	UserName string
	NickName string
}

var (
	ErrMissingLoginValues = errors.New("用户名或密码不能为空")

	ErrFailedAuthentication = errors.New("用户名或密码错误")
)

func JwtMiddleware() (authMiddleware *jwt.GinJWTMiddleware, err error) {
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			user, err := httpService.FindUser(userID)
			if err != nil {
				return nil, ErrFailedAuthentication
			}

			encrypt := utils.PasswordEncrypt(password, userID)
			if encrypt != user.Password.String {
				return nil, ErrFailedAuthentication
			}

			return &User{
				Id:       user.Id,
				UserName: user.Username.String,
				NickName: user.Nickname.String,
			}, nil
		},
		LoginResponse: func(c *gin.Context, status int, tokenString string, expire time.Time) {
			ctx := context.Background()
			token, _ := authMiddleware.ParseTokenString(tokenString)
			claims := jwt.ExtractClaimsFromToken(token)
			UserName := claims[identityKey].(string)

			db.RedisClient.Set(ctx, "token:"+UserName, tokenString, authMiddleware.Timeout)

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"token":   tokenString,
				"expire":  expire.Format(time.RFC3339),
				"message": "登录成功",
			})
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			token := jwt.GetToken(c)

			user := data.(*User)
			ctx := context.Background()
			userToken := db.RedisClient.Get(ctx, "token:"+user.UserName)
			log.Println(userToken)
			if userToken.Val() == token {
				return true
			}

			return false
			//if v, ok := data.(*User); ok && v.UserName == "root" {
			//	return true
			//}
			//return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
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
		TokenLookup: "header: Authorization, query: token, cookie: jwt_middleware",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	return
}
