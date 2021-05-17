package classes

import (
	"fmt"
	"github.com/fyxemmmm/chitanda/chitanda"
	"github.com/fyxemmmm/chitanda/src/models"
	"github.com/gin-gonic/gin"
)

type UserClass struct {
	*chitanda.SqlXAdapter
}

func NewUserClass() *UserClass {
	return &UserClass{}
}


func (this *UserClass) UserTest(ctx *gin.Context) string {
	return "测试"
}


func (this *UserClass) UserList(ctx *gin.Context) chitanda.Models {
	users := []*models.UserModel{
		{UserId: 101, UserName: "feixiang101"},
		{UserId: 102, UserName: "feixiang102"},
	}
	return chitanda.ToModels(users)
}

func (this *UserClass) UserDetail(ctx *gin.Context) chitanda.Model {
	user := &models.UserModel{}
	err := ctx.BindUri(user)
	chitanda.Error(err, "用户id参数不合法")
	fmt.Println(user.UserId)

	sql := "select id, name, age, email from my.user where id = 2"
	err = this.GetContext(ctx, user, sql)
	if err != nil {
		chitanda.Error(err)
	}
	return user
}


func (this *UserClass) Build(chitanda *chitanda.Chitanda)  {
	chitanda.Handle("GET", "/test", this.UserTest)
	chitanda.Handle("GET", "/user/:id", this.UserDetail)
	chitanda.Handle("GET", "/user-list", this.UserList)
}
