package chapter02

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

type Student struct {
	Id       int
	Name     string `form:"user_name"`
	Password string `form:"password"`
}

func (u *UserController) Get() {
	//第一种方式
	//id := u.GetString(":id")
	//第二种
	//id := u.GetString("id")
	// id, err := u.GetInt("id")
	// if err != nil {
	// 	u.Ctx.WriteString("只能传int类型的参数")
	// }
	// u.Data["id"] = id
	u.TplName = "chapter02/user.html"
}

func (u *UserController) Post() {
	//获取post
	// user_name := u.GetString("user_name")
	// password := u.GetString("password")
	// fmt.Println("用户名：" + user_name)
	// fmt.Println("密码：" + password)

	//form表单解析到结构体
	student := Student{}
	u.ParseForm(&student)

	fmt.Println("用户名：" + student.Name)
	fmt.Println("密码：" + student.Password)
	u.Ctx.WriteString("提交成功！")
}
