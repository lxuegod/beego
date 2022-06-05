package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	//渲染字符串到浏览器
	c.Ctx.WriteString("hello world!")

	//结构体渲染
	// var user User
	// user.Id = 1
	// user.Name = "hallen"
	// user.Age = 18
	user := User{
		Id:   2,
		Name: "hallen",
		Age:  18,
	}

	c.Data["user"] = user

	//数组渲染
	arr := [5]int{5, 4, 3, 2, 1}
	c.Data["arr"] = arr

	//结构体+数组渲染
	users := [3]User{
		{
			Id:   1,
			Name: "hallen1",
			Age:  18,
		},
		{
			Id:   2,
			Name: "hallen2",
			Age:  19,
		},
		{
			Id:   3,
			Name: "hallen3",
			Age:  20,
		},
	}

	c.Data["users"] = users

	//map渲染
	// map_data := map[string]string{
	// 	"name": "hallen",
	// 	"age":  "18",
	// }
	// c.Data["map_data"] = map_data

	map_data_inter := map[string]interface{}{
		"name": "hallen",
		"age":  18,
	}
	c.Data["map_data"] = map_data_inter

	//map+结构体渲染
	map_struct := map[string]User{
		"user1": {
			Id:   1,
			Name: "hallen1",
			Age:  18,
		},
		"user2": {
			Id:   2,
			Name: "hallen2",
			Age:  19,
		},
	}
	c.Data["map_struct"] = map_struct

	//切片渲染
	slice := []int{9, 8, 7, 6, 5, 4, 3}
	c.Data["slice"] = slice

	c.TplName = "index.html"
}
