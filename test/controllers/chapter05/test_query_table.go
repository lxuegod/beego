package chapter05

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"test/models/chapter05"
)

type TestQueryTable struct {
	beego.Controller
}

func (t *TestQueryTable) Get() {

	o := orm.NewOrm()
	user := chapter05.User{}
	users := []chapter05.User{}
	qs := o.QueryTable(new(chapter05.User))
	//one
	qs.One(&user)

	//filte	包含条件 多个filter用and连接
	qs.Filter("name", "hallenx").One(&user)

	//all	所有数据，返回对应的结果集对象
	//qs.All(&users)

	//exclude	排除条件
	//qs.Exclude("name", "hallen").All(&users)

	//limit限制最大返回数据行数 offset设置偏移行数，意思是从什么位置开始查询
	//qs.Limit(5).Offset(1).All(&users)

	//GroupBy	分组
	//qs.GroupBy("age").All(&users)

	//orderBy 默认升序 降序-
	//qs.OrderBy("-age").All(&users)

	//distinct	去重
	//qs.Distinct().All(&users, "age", "address")

	//count	统计个数，没有参数，后面不能跟查询条数的方法
	//count, _ := qs.Count()
	count, _ := qs.Filter("age", 18).Count()

	//exit
	exit := qs.Filter("age", 20).Exist()

	//update
	//
	//qs.Filter("age", 18).Update(orm.Params{
	//	"address": "rrrr",
	//})

	//delete
	//qs.Filter("age", 11).Delete()

	//insert	不是QueryTable的接口
	user2 := chapter05.User{}
	user2.Name = "zhiliao"
	profile := chapter05.Profile{}
	o.QueryTable(new(chapter05.Profile)).Filter("id", 4).One(&profile)
	fmt.Println("a====================")
	fmt.Println(profile)
	user2.Profile = &profile
	_, err := o.Insert(&user2)
	fmt.Println(err)

	t.Data["user"] = user
	t.Data["users"] = users
	t.Data["count"] = count
	t.Data["exit"] = exit

	t.TplName = "chapter05/test_query_table.html"
}
