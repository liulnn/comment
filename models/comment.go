package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

func init() {
	orm.RegisterDriver("postgres", orm.DR_Postgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=comment sslmode=disable")
	orm.RegisterModel(new(Comment))
}

type Comment struct {
	Id        int64
	Topic     string
	Source    string
	Username  string
	Content   string
	CreatedAt time.Time
}

func (c *Comment) TableName() string {
	return "comment"
}

func AddComment(topicId, source, username, content string) (id int64, err error) {
	o := orm.NewOrm()
	return o.Insert(&Comment{
		Topic:    topicId,
		Source:   source,
		Username: username,
		Content:  content,
	})
}

func GetComments(topicId string, pageSize int, pageNum int) (comments []*Comment, err error) {
	o := orm.NewOrm()
	c := new(Comment)
	qs := o.QueryTable(c)
	_, err = qs.Filter("topic", topicId).Offset(pageSize * (pageNum - 1)).Limit(pageSize).All(&comments)
	return comments, err
}

func GetComment(topicId string, id int64) (c *Comment, err error) {
	o := orm.NewOrm()
	c = &Comment{Id: id}
	err = o.Read(c)
	return c, err
}

func DeleteComment(topicId string, id int64) (success bool, err error) {
	o := orm.NewOrm()
	num, err := o.Delete(&Comment{Id: id})
	if num > 0 {
		success = true
	}
	success = false
	return success, err
}
