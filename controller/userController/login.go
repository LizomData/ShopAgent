package userController

import (
	"ShopAgent/model/requestModel"
	"ShopAgent/util"
	"ShopAgent/util/database/userDbService"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登陆
// @Accept       json
// @Produce      json
// @Tags 用户管理
// @Param  body body model.User true "登录凭证"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/account/login [post]
func Login(c *gin.Context) {
	vail, user := validateForm(c)
	if !vail {
		return
	}
	var err error
	err, user = userDbService.Instance.CheckPassword(user)
	//密码校验
	if err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.LoginFailed, "用户名或密码错误"))
		return
	}

	// 生成JWT
	tokenString, err := util.GenerateToken(user, 2400)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.TokenGenerationFailed, "生成token失败:"+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(gin.H{
		"userInfo": gin.H{"username": user.Username},
		"token":    tokenString,
	}))

}
