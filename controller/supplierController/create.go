package supplierController

import (
	"ShopAgent/common"
	"ShopAgent/model"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/supplierDbService"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var req model.Supplier
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(requestModel.ResponseFailure(requestModel.ParameterError, "解析参数失败: "+err.Error()))
		return
	}
	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户id失败: "+err.Error()))
		return
	}
	req.UserID = user.ID
	supplier, err := supplierDbService.Instance.Create(&req)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "创建失败: "+err.Error()))
		return
	}
	c.JSON(requestModel.ResponseSuccess(supplier))
}
