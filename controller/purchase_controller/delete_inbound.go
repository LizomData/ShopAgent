package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
)

// DeleteInbound 删除进货单
func DeleteInbound(c *gin.Context) {
	var deleteInboundDTO dto.DeleteInboundDTO
	if err := c.ShouldBind(&deleteInboundDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	err = purchaseDbService.Instance.DeleteInbound(user.ID, deleteInboundDTO.InboundID)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "删除进货单失败: "+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(nil))
}
