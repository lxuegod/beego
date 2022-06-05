package cms

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"hello/models"
	"hello/utils"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "cms/login.html"
}
func (l *LoginController) Post() {
	username := l.GetString("username") //特殊字符过滤。?/
	password := l.GetString("password") //长度校验

	md5_pwd := utils.GetMd5(password)
	o := orm.NewOrm()
	exist := o.QueryTable(new(models.User)).Filter("user_name", username).Filter("password", md5_pwd).Exist()

	if exist {
		l.SetSession("cms_user_name", username)
		l.Redirect(beego.URLFor("MainController.Get"), 302)
		fmt.Println("登录成功")

	} else {
		l.Redirect(beego.URLFor("LoginController.Get"), 302)
	}

}
