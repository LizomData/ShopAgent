package userController

import (
	"ShopAgent/model/requestModel"
	"ShopAgent/util"
	"ShopAgent/util/database/userDbService"
	"github.com/gin-gonic/gin"
	"time"
)

// @Summary 用户注册
// @Accept       json
// @Produce      json
// @Tags 用户管理
// @Param  body body model.User true "注册凭证"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/account/register [post]
func Register(c *gin.Context) {

	vail, user := validateForm(c)
	if !vail {
		return
	}

	//查询重复
	if err, _ := userDbService.Instance.GetUserByUsername(user.Username); err == nil {
		c.JSON(requestModel.ResponseFailure(requestModel.RegisterAlready, "用户已被注册"))
		return
	}

	//创建用户
	if err := userDbService.Instance.CreateUser(user); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.RegisterFailed, "注册失败"))
		return
	}

	err, user_new := userDbService.Instance.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.RegisterFailed, "注册失败"))
		return
	}
	// 生成JWT
	tokenString, err := util.GenerateToken(user_new, 240)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.TokenGenerationFailed, "生成token失败"))
		return
	}

	c.JSON(requestModel.ResponseSuccess(gin.H{"username": user.Username, "token": tokenString, "timeStamp": time.Now().Unix()}))
}
