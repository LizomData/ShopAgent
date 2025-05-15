package requestModel

import "github.com/gin-gonic/gin"

func ResponseFailure(error_code int, msg string) (int, gin.H) {
	return 200, gin.H{
		"code": error_code,
		"msg":  msg,
		"data": gin.H{},
	}
}
func ResponseSuccess(data any) (int, gin.H) {
	return 200, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": data,
	}
}

const (
	LoginStatusInvalid int = iota + 1001
	NotLoggedIn
	InvalidTokens
	LoginFailed
	ParameterError
	TokenGenerationFailed
	RegisterFailed
	RegisterAlready
	IllegalCharacter
	IncorrectFormat
	NotUser
	NotPrivileged
	NotFile
	RepoFailed
	InvalidType
	SBOMFailed
	FileUnzipFailed
	FileNotFound
	UploadFailed
	NotSbomReport
	NotFoundReport
)
const (
	Success int = iota
)
