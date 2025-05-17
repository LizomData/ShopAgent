package util

import (
	"ShopAgent/model"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/userDbService"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWT密钥
var jwtSecret = []byte("2004qwe")

// JWT Claims
type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(user model.User, expiredTime_hour time.Duration) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * expiredTime_hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(requestModel.ResponseFailure(requestModel.LoginStatusInvalid, "authHeader为空"))
			c.Abort()
			return
		}
		//tokenString := authHeader[len("Bearer "):]
		tokenString := authHeader
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(requestModel.ResponseFailure(requestModel.InvalidTokens, "无效token:"+err.Error()))
			c.Abort()
			return
		}

		err, user := userDbService.Instance.GetUserById(claims.UserID)
		if err != nil {
			c.JSON(requestModel.ResponseFailure(requestModel.NotUser, "查询用户出现错误:"+err.Error()))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
