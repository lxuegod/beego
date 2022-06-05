package chapter04

import (
	"fmt"
	"github.com/astaxie/beego/validation"

	beego "github.com/beego/beego/v2/server/web"
)

type ValidControlle struct {
	beego.Controller
}

type MyUser struct { //映射
	Id       int    `form:"-"`
	UserName string `form:"user_name" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

func (v *ValidControlle) Get() {

	v.TplName = "chapter04/test_valid.html"
}
func (v *ValidControlle) Post() {
	var user MyUser
	v.ParseForm(&user)
	fmt.Println("===============")
	fmt.Println(user)

	valid := validation.Validation{}

	var MessageTmpls = map[string]string{
		"Required": "不能为空  ",
		"Min":      "不能小于 %d",
		"Max":      "不能大于 %d",
		"Range":    "介于 %d to %d",
	}

	validation.SetDefaultMessage(MessageTmpls)

	b, err := valid.Valid(&user)

	if err != nil {
		v.Ctx.WriteString("结构体上的tag有问题")
	}
	if !b {
		for _, err_msg := range valid.Errors {
			fmt.Println(err_msg.Key)
			fmt.Println((err_msg.Message))
			v.Ctx.WriteString(err_msg.Message)
			return
		}
	}

	v.Ctx.WriteString("提交了!")
}
