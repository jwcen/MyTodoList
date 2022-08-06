package service

import (
	"MyTodoList/model"
	"MyTodoList/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status int `json:"status" form:"status"`  // 0 已做  1 未做
}

// Create 创建备忘录逻辑
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	// 先找到user
	model.DB.First(&user, id)
	// 用户创建备忘录
	task := model.Task{
		User: user,
		Uid: user.ID,
		Title: service.Title,
		Content: service.Content,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "成功创建一条备忘录",
	}
}