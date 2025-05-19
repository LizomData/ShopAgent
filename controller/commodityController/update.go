package commodityController

import (
	"ShopAgent/common"
	"ShopAgent/model"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/commodityDbService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var req model.CommodityInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(req.ID)
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}
	if _, err := commodityDbService.Instance.Update(user.ID, &req); err != nil {
		c.JSON(requestModel.ResponseFailure(999, "更新商品信息失败: "+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(req))

}
