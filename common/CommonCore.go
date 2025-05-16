package common

import (
	"ShopAgent/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (model.User, error) {
	// 1. 检查是否存在用户对象
	rawUser, exists := c.Get("user")
	if !exists {
		return model.User{}, fmt.Errorf("user not found in context")
	}

	// 2. 安全的类型断言
	user, ok := rawUser.(model.User)
	if !ok {
		return model.User{}, fmt.Errorf("invalid user type in context")
	}

	return user, nil
}
