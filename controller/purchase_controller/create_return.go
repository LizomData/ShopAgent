package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateReturn 创建退货单
func CreateReturn(c *gin.Context) {
	var purchaseReturnDTO dto.PurchaseReturnDTO
	if err := c.ShouldBindJSON(&purchaseReturnDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	returnVO, err := purchaseDbService.Instance.CreateReturn(&purchaseReturnDTO, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  returnVO,
		})
		return
	}

	response := gin.H{
		"code":    200,
		"message": "退货单创建成功",
		"error":   "",
		"data":    returnVO,
	}

	c.JSON(http.StatusOK, response)
}
