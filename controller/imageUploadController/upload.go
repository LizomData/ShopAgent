package imageUploadController

import (
	"ShopAgent/common"
	"ShopAgent/model/requestModel"
	"ShopAgent/util/database/imageUploadDbService"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "获取用户id失败:"+err.Error()))
		return
	}

	// 获取上传文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "请选择上传文件"))
		return
	}

	// 调用上传模块
	record, err := Uploader.UploadFile(fileHeader)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "上传文件失败:"+err.Error()))
		return
	}

	//保存上传记录
	record, err = imageUploadDbService.Instance.Create(user.ID, record)
	if err != nil {
		c.JSON(requestModel.ResponseFailure(999, "保存上传文件记录失败:"+err.Error()))
		return
	}

	c.JSON(requestModel.ResponseSuccess(record))
}
