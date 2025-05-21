package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
)

// UpdateReturn 修改退货单
func UpdateReturn(c *gin.Context) {
	var updateReturnDTO dto.UpdateReturnDTO
	if err := c.ShouldBind(&updateReturnDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	if err := purchaseDbService.Instance.UpdateReturn(user.ID, &updateReturnDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(999, "修改退货单失败: "+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(nil))
}
