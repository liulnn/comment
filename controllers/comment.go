package controllers

import (
	"bytes"
	"comment/models"
	"crypto/md5"
	"darling"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type CommentsCtrl struct {
	darling.Controller
	AccessController
}

func (c *CommentsCtrl) Prepare() {
	c.Access(c.Request)
}

func (c *CommentsCtrl) Post() {
	content := "this is content"
	if c.AccountId <= 0 {
		c.Response.StatusCode = http.StatusForbidden
		return
	}
	comment, err := models.AddComment(c.PathParams[0], c.AccountId, content)
	if err != nil {
		c.Response.StatusCode = http.StatusInternalServerError
		return
	}
	data, _ := json.Marshal(comment)
	h := md5.New()
	io.WriteString(h, string(data))
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
	c.Response.StatusCode = http.StatusCreated
	c.Response.Header["ETag"] = buffer.String()
	c.Response.ContentType = "application/json; charset=utf-8"
	c.Response.Content = data
}

func (c *CommentsCtrl) Get() {
	comments, err := models.GetComments(c.PathParams[0], 10, 0)
	if err != nil {
		c.Response.StatusCode = http.StatusInternalServerError
		return
	}
	data, _ := json.Marshal(comments)
	c.Response.ContentType = "application/json; charset=utf-8"
	c.Response.Content = data
}

type CommentCtrl struct {
	AccessController
	darling.Controller
}

func (c *CommentCtrl) Prepare() {
	c.Access(c.Request)
}

func (c *CommentCtrl) Get() {
	topicId := c.PathParams[0]
	commentId, _ := strconv.ParseInt(c.PathParams[1], 10, 64)
	comment, err := models.GetComment(topicId, commentId)
	if err != nil {
		c.Response.StatusCode = http.StatusInternalServerError
		return
	}
	data, _ := json.Marshal(comment)
	h := md5.New()
	io.WriteString(h, string(data))
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
	c.Response.Header["ETag"] = buffer.String()
	c.Response.ContentType = "application/json; charset=utf-8"
	c.Response.Content = data
}
func (c *CommentCtrl) Delete() {
	if c.AccountId <= 0 {
		c.Response.StatusCode = http.StatusForbidden
		return
	}
	topicId := c.PathParams[0]
	commentId, _ := strconv.ParseInt(c.PathParams[1], 10, 64)
	err := models.DeleteComment(topicId, commentId, c.AccountId)
	if err != nil {
		c.Response.StatusCode = http.StatusInternalServerError
		return
	}
	c.Response.StatusCode = http.StatusNoContent
}
