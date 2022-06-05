package chapter02

import (
	beego "github.com/beego/beego/v2/server/web"
)

type OtherController struct {
	beego.Controller
}

type Teacher struct {
	Id   int
	Name string
	Age  int
}

func (o *OtherController) Get() {
	// teacher := Teacher{Id: 1, Name: "hallen", Age: 18}
	// //json格式
	// // o.Data["json"] = teacher
	// // o.ServeJSON()

	// //xml格式
	// // o.Data["xml"] = teacher
	// // o.ServeXML()

	// //yaml格式
	// o.Data["yaml"] = teacher
	// o.ServeYAML()
	o.TplName = "chapter02/other.html"
}

func (o *OtherController) Post() {
	teacher := Teacher{Id: 1, Name: "hallen", Age: 18}
	o.Data["yaml"] = teacher
	o.ServeYAML()
}
