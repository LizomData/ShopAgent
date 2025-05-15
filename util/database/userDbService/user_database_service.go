package userDbService

import (
	"ShopAgent/model"
	"ShopAgent/util/database"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
)

var Instance = GetInstance()

func GetInstance() *UserDbService {
	// 初始化节点（确保每个服务实例的 nodeID 唯一）
	nodeID := int64(1)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		fmt.Println("初始化雪花节点失败: " + err.Error())
		return nil
	}
	return &UserDbService{node}
}

type UserDbService struct {
	node *snowflake.Node
}

// 创建用户
func (u *UserDbService) CreateUser(user model.User) error {
	if user.ID == 0 { // 如果未手动设置 ID，则自动生成
		user.ID = u.GenerateSnowflakeID()
	}
	err, hash := u.GetPasswordHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	result := database.GormDB.Create(&user)
	return result.Error
}

// 查询用户
func (u *UserDbService) GetUserByUsername(username string) (error, model.User) {
	var user model.User
	result := database.GormDB.First(&user, "username = ?", username)
	return result.Error, user
}

// 查询用户
func (u *UserDbService) GetUserById(id int64) (error, model.User) {
	var user model.User
	result := database.GormDB.First(&user, "id = ?", id)
	return result.Error, user
}

// 设置密码（生成哈希）
func (u *UserDbService) GetPasswordHash(password string) (error, string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return err, string(hashedBytes)
}

// 验证密码
func (u *UserDbService) CheckPassword(user model.User) (error, model.User) {
	err, _user := u.GetUserByUsername(user.Username)
	if err != nil {
		return err, model.User{}
	}
	return bcrypt.CompareHashAndPassword([]byte(_user.Password), []byte(user.Password)), _user
}

// GenerateSnowflakeID 生成雪花 ID
func (u *UserDbService) GenerateSnowflakeID() int64 {
	return u.node.Generate().Int64()
}
