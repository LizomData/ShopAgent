package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
)

// UpdateInbound 修改进货单
func UpdateInbound(c *gin.Context) {
	var updateInboundDTO dto.UpdateInboundDTO
	if err := c.ShouldBind(&updateInboundDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	if err := purchaseDbService.Instance.UpdateInbound(user.ID, &updateInboundDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(999, "修改进货单失败: "+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(nil))
}
