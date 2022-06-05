package utils

import (
	"github.com/astaxie/beego/context"
	beego "github.com/beego/beego/v2/server/web"
)

func CmsLoginFilter(ctx *context.Context) {
	cms_user_name := ctx.Input.Session("cms_user_name")

	if cms_user_name == nil {
		ctx.Redirect(302, beego.URLFor("LoginController.Get"))
	}
}
