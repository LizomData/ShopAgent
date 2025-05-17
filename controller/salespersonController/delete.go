package salespersonController

import (
	"ShopAgent/common"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/salespersonDbService"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}
	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户id失败: "+err.Error()))
		return
	}
	if err = salespersonDbService.Instance.Delete(user.ID, req.ID); err != nil {
		c.JSON(requestModel.ResponseFailure(999, "删除失败: "+err.Error()))
		return
	}
	c.JSON(requestModel.ResponseSuccess(gin.H{}))

}
