package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/dto"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateInbound 创建进货单
func CreateInbound(c *gin.Context) {
	var purchaseInboundDTO dto.PurchaseInboundDTO
	if err := c.ShouldBindJSON(&purchaseInboundDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}

	inboundVO, err := purchaseDbService.Instance.CreateInbound(&purchaseInboundDTO, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	response := gin.H{
		"code":    200,
		"message": "进货单创建成功",
		"data":    inboundVO,
	}

	c.JSON(http.StatusOK, response)
}
