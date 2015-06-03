package models

import (
	"errors"
	"strconv"
	"time"
)

func init() {

}

type Comment struct {
	Id        int64
	Topic     string
	Source    string
	Username  string
	Content   string
	CreatedAt time.Time
}

func AddComment(c Comment) string {
	return c.Id
}

func GetComment(id int64) (c *Comment, err error) {

}

func DeleteComment(id int64) (err error) {

}

func GetComments(pageSize int, pageNum int) map[int64]*Comment {

}
