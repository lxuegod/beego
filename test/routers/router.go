package routers

import (
	"test/controllers/chapter02"
	"test/controllers/chapter03"
	"test/controllers/chapter04"
	"test/controllers/chapter05"
	"test/controllers/chapter06"

	"test/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
	//第二种  url为/user?id=参数
	beego.Router("/user", &chapter02.UserController{})
	//第一种
	//beego.Router("/user/?:id:int", &chapter02.UserController{})
	beego.Router("/add_user", &chapter02.UserController{})
	beego.Router("/upload", &chapter02.UploadController{})

	beego.Router("/other", &chapter02.OtherController{})
	beego.Router("/chap01", &chapter03.Chap01rController{})
	//自定义路由
	//beego.Router("/chap01", &chapter03.Chap01rController{}, "get:方法")

	// //测试过滤器的url
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/test_valid", &chapter04.ValidControlle{})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &chapter04.LoginController{})

	//需要登录的
	beego.Router("/cms/test1", &chapter04.LoginController{})
	beego.Router("/cms/test2", &chapter04.LoginController{})
	beego.Router("/cms/test3", &chapter04.LoginController{})
	beego.Router("/cms/test4", &chapter04.LoginController{})

	beego.Router("/test_exper", &chapter05.ExperController{})

	beego.Router("/test_query_table", &chapter05.TestQueryTable{})
	beego.Router("/test_logs", &chapter06.LogsController{})

}
