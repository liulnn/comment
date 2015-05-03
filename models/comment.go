package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"time"
)

var engine *xorm.Engine

func init() {
	engine, _ = xorm.NewEngine("postgres", "postgres://postgres:123456@localhost/comment?sslmode=disable")
	engine.Sync2(new(Account), new(Comment))
}

type Account struct {
	Id          int64  `xorm:"pk autoincr"`
	AppSource   string `xorm:"varchar(16) index(app_token)"`
	AccessToken string `xorm:"varchar(32) index(app_token)"`
	Username    string `xorm:"varchar(100)"`
}

func (a Account) TableName() string {
	return "account"
}

func AddAccount(appSource string, accessToken string, username string) (err error) {
	account := &Account{AppSource: appSource, AccessToken: accessToken, Username: username}
	_, err = engine.Insert(account)
	return err
}

func GetAccount(appSource string, accessToken string) (accountId int64, err error) {
	account := &Account{AppSource: appSource, AccessToken: accessToken}
	_, err = engine.Get(account)
	return account.Id, err
}

type Comment struct {
	AccountId int64     `xorm: bigint`
	Id        int64     `xorm: "pk autoincr"`
	TopicId   string    `xorm:"varchar(32) index"`
	Content   string    `xorm:"varchar(255)"`
	CreatedAt time.Time `xorm:"datetime created"`
}

func (c Comment) TableName() string {
	return "comment"
}

func AddComment(topicId string, accountId int64, content string) (comment *Comment, err error) {
	comment = &Comment{AccountId: accountId, TopicId: topicId, Content: content}

	_, err = engine.Insert(comment)
	return comment, err
}

func GetComments(topicId string, pageSize int, pageNum int) (comments []Comment, err error) {
	comments = make([]Comment, 0)
	err = engine.Where("topic_id = ?", topicId).Limit(pageSize, pageSize*pageNum).Find(&comments)
	return comments, err
}

func GetComment(topicId string, commentId int64) (comment *Comment, err error) {
	comment = &Comment{}
	success, err := engine.Id(commentId).Get(comment)
	if success {
		return comment, nil
	}
	return nil, err
}
func DeleteComment(topicId string, commentId int64, accountId int64) error {
	_, err := engine.Id(commentId).Delete(new(Comment))
	return err
}
