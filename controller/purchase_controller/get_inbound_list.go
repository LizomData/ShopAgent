package purchase_controller

import (
	"ShopAgent/common"
	"ShopAgent/model/requestModel"
	purchaseDbService "ShopAgent/util/database/purchase_db_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetInboundList 获取进货单列表
func GetInboundList(c *gin.Context) {
	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户失败: "+err.Error()))
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	list, total, err := purchaseDbService.Instance.GetInboundList(page, pageSize, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"data":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
