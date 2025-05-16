package commodityController

import (
	"ShopAgent/common"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/commodityDbService"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}
	_commodity_info, err := commodityDbService.Instance.GetById(user.ID, req.ID)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "查询不到该商品: "+err.Error()))
		return
	}

	_commodity_info.Name = req.Name
	_commodity_info.Type = req.Type
	_commodity_info.Specifications = req.Specifications
	_commodity_info.Quantity = req.Quantity
	_commodity_info.Price = req.Price
	_commodity_info.CartImage = req.CartImage

	if _, err := commodityDbService.Instance.Update(&_commodity_info); err != nil {
		c.JSON(requestModel.ResponseFailure(999, "更新商品信息失败: "+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(_commodity_info))

}
