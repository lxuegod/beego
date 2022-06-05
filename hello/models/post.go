package models

import "time"

//type Tag struct {
//	Id   int
//	Name string
//}

type Post struct {
	Id         int       `orm:"pl;auto"`
	Title      string    `orm:"description(帖子标题)"`
	Desc       string    `orm:"description(帖子描述)"`
	Content    string    `orm:"size(4000);description(帖子内容)"`
	Cover      string    `orm:"description(帖子封面图);default(static/upload/no_p.jpg)"`
	ReadNum    int       `orm:"description(帖子阅读数);default(0)"`
	StarNum    int       `orm:"description(帖子点赞数);default(0)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`

	Author *User `orm:"description(帖子作者);rel(fk)"`
}

func (p *Post) TableName() string {
	return "sys_post"
}
