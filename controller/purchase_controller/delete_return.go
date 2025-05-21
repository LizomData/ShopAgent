package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
)

// DeleteReturn 删除退货单
func DeleteReturn(c *gin.Context) {
	var deleteReturnDTO dto.DeleteReturnDTO
	if err := c.ShouldBind(&deleteReturnDTO); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	err = purchaseDbService.Instance.DeleteReturn(user.ID, deleteReturnDTO.ReturnID)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "删除退货单失败: "+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(nil))
}
