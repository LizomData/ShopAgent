package salespersonController

import (
	"ShopAgent/common"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/salespersonDbService"
	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	var req QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}
	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户id失败: "+err.Error()))
		return
	}
	salespersons, err := salespersonDbService.Instance.Query(user.ID, *req.Page, *req.Size)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "查询失败:"+err.Error()))
		return
	}
	c.JSON(requestModel.ResponseSuccess(salespersons))
}
