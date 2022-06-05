package chapter05

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int    `orm:"pk;auto"`
	Name       string `orm:"index;unique"` //索引
	Age        int
	Addr       string    `orm:"null;column(address) "`
	Dec        string    `orm:"size(2000"`
	Price      float64   `orm:"digits(12);decimals(4)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	//1启用， 0停用
	Status   int        `orm:"default(1);description(1启用,0停用)"` //default默认值,description加注释
	XXX      string     `orm:"-"`
	Profile  *Profile   `orm:"rel(one)"`      //一对一
	Articles []*Article `orm:"reverse(many)"` //一对多
}

type Profile struct {
	Id     int
	IdCard string
	Cover  string
	User   *User `orm:"reverse(one)""`
}

type Article struct {
	Id      int
	Title   string
	Content string `orm:"size(2000)"` //一对一
	User    *User  `orm:"rel(fk)"`    //一对多
}

type Post struct {
	Id   int
	Tile string
	Tags []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

//重写表名
func (u *User) TableName() string { //正向关系，有外键字段
	return "sys_user"
}

func (p *Profile) TableName() string { //反向关系，只是关系，没有外键
	return "sys_profile"
}

func (a *Article) TableName() string { //正向关系，有外键，外键在多的一方，user模型  反向关系，没有外键
	return "sys_article"
}

func (a *Post) TableName() string { //正向关系，没有外键
	return "sys_post"
}

func (a *Tag) TableName() string { //反向关系，没有外键
	return "sys_tag"
}

func init() {
	orm.RegisterModel(new(User), new(Profile), new(Article), new(Post), new(Tag))
	
	orm.Debug = true
}
