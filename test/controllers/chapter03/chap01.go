package chapter03

import (
	beego "github.com/beego/beego/v2/server/web"
)

type Chap01rController struct {
	beego.Controller
}

func (c *Chap01rController) Get() {
	c.TplName = "chapter03/chap01.html"
}
