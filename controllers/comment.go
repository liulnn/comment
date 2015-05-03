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
	AccessController
	darling.Controller
}

func (c *CommentsCtrl) Post() {
	accountId, err := c.Access(c.Request)
	if err != nil {
		c.Response.WriteHeader(http.StatusForbidden)
	}
	content := "this is content"
	comment, err := models.AddComment(c.PathParams[0], accountId, content)
	if err != nil {
		c.Response.WriteHeader(http.StatusInternalServerError)
	}
	data, _ := json.Marshal(comment)
	c.Response.Header().Add("Content-Type", "application/json; charset=utf-8")
	h := md5.New()
	io.WriteString(h, string(data))
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
	c.Response.Header().Add("ETag", buffer.String())
	io.WriteString(c.Response, string(data))
}

func (c *CommentsCtrl) Get() {
	comments, err := models.GetComments(c.PathParams[0], 10, 0)
	if err != nil {
		c.Response.WriteHeader(http.StatusInternalServerError)
	}
	data, _ := json.Marshal(comments)
	c.Response.Header().Add("Content-Type", "application/json; charset=utf-8")
	io.WriteString(c.Response, string(data))
}

type CommentCtrl struct {
	AccessController
	darling.Controller
}

func (c *CommentCtrl) Get() {
	topicId := c.PathParams[0]
	commentId, _ := strconv.ParseInt(c.PathParams[1], 10, 64)
	comment, err := models.GetComment(topicId, commentId)
	if err != nil {
		c.Response.WriteHeader(http.StatusInternalServerError)
	}
	data, _ := json.Marshal(comment)
	c.Response.Header().Add("Content-Type", "application/json; charset=utf-8")
	h := md5.New()
	io.WriteString(h, string(data))
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
	c.Response.Header().Add("ETag", buffer.String())
	io.WriteString(c.Response, string(data))

}
func (c *CommentCtrl) Delete() {
	topicId := c.PathParams[0]
	commentId, _ := strconv.ParseInt(c.PathParams[1], 10, 64)
	accountId, err := c.Access(c.Request)
	if err != nil {
		c.Response.WriteHeader(http.StatusForbidden)
	}
	err = models.DeleteComment(topicId, commentId, accountId)
	if err != nil {
		c.Response.WriteHeader(http.StatusInternalServerError)
	}
	c.Response.WriteHeader(http.StatusNoContent)

}
