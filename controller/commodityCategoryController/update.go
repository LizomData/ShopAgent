package commodityCategoryController

import (
	"ShopAgent/common"
	"ShopAgent/model"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/commodityCategoryDbService"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var req model.CommodityCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}
	if _, err := commodityCategoryDbService.Instance.Update(user.ID, &req); err != nil {
		c.JSON(requestModel.ResponseFailure(999, "更新失败: "+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(req))

}
