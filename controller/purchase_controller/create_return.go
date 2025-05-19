package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
)

// CreateReturn 创建退货单
func CreateReturn(c *gin.Context) {
	var purchaseReturnDTO dto.PurchaseReturnDTO
	if err := c.ShouldBind(&purchaseReturnDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	returnVO, err := purchaseDbService.Instance.CreateReturn(&purchaseReturnDTO, user.ID)

	c.JSON(requestModel.ResponseSuccess(returnVO))
}
