package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
)

// CreateInbound 创建进货单
func CreateInbound(c *gin.Context) {
	var purchaseInboundDTO dto.PurchaseInboundDTO
	if err := c.ShouldBind(&purchaseInboundDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	inboundVO, err := purchaseDbService.Instance.CreateInbound(&purchaseInboundDTO, user.ID)
	if err != nil {

		c.JSON(requestModel.ResponseFailure(999, "创建进货订单失败: "+err.Error()))

		return
	}

	c.JSON(requestModel.ResponseSuccess(inboundVO))
}
