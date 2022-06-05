package chapter05

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"test/models/chapter05"
)

type ExperController struct {
	beego.Controller
}

func (e *ExperController) Get() {
	o := orm.NewOrm()
	//o.QueryTable("sys_user")  第一种方式
	qs := o.QueryTable(new(chapter05.User)) //模型方式
	user := chapter05.User{}
	//exact
	//qs.Filter("name", "hallen").One(&user)
	//contains
	//qs.Filter("name__contains", "lle").One(&user)
	//gt/gte
	//qs.Filter("age__gt", "19").One(&user)
	//lt/lte
	//qs.Filter("age__lt", "19").One(&user)

	//startswith	以什么开头
	//qs.Filter("name__startswith", "ha").One(&user)
	//endswith 以什么结尾
	//qs.Filter("name__endswith", "en").One(&user)
	//in 在什么范围内
	//qs.Filter("age__in", "15,16,27,19,18").One(&user)
	//isnull,true查询为null的
	qs.Filter("address__isnull", true).One(&user)

	fmt.Println(user)
	e.Data["user"] = user
	e.TplName = "chapter05/tets_exper.html"
}
