package service

import (
	"MyTodoList/model"
	"MyTodoList/pkg/utils"
	"MyTodoList/serializer"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

// Register 用户注册逻辑
func (s *UserService) Register() serializer.Response {
	var user model.User
	var count int
	// 查询数据库是否存在，不存在就创建用户
	model.DB.Model(&model.User{}).Where("user_name=?", s.UserName).
		First(&user).Count(&count)
	
	if count == 1 {
		return serializer.Response{
			Status: 403,
			Msg: "此用户已存在",
		}
	}

	// 否则，创建用户
	user.UserName = s.UserName
	// 密码存到数据库前要进行加密
	if err := user.SetPassword(s.Password); err != nil {
		return serializer.Response{
			Status: 403,
			Msg: "密码错误",
		}
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "数据库操作错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg: "用户注册成功",
	}
}

// Login 用户登录逻辑
func (s *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", s.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 403,
				Msg: "用户不存在",
			}
		}
		// 其他错误
		return serializer.Response{
			Status: 500,
			Msg: "服务器内部错误",
		}
	}

	// 验证密码
	if !user.CheckPassword(s.Password) {
		return serializer.Response{
			Status: 403,
			Msg: "密码错误", 
		}
	}

	// 发送token给前端存储，以便其他需要验证身份的功能
	// 如，创建一个备忘录就需要token，否则不知道备忘录由谁创建的
	token, err := utils.GenerateToken(user.ID, s.UserName, s.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "Token签发错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User: serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}