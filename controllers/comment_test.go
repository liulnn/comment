package controllers

import (
	"bytes"
	"crypto/md5"
	"darling"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestCommentsCtrlPost(t *testing.T) {
	var topicId string = "1"
	req, _ := http.NewRequest("POST", fmt.Sprintf("/topics/%s/comments", topicId), nil)
	resp := &darling.Response{Header: make(map[string]string)}
	c := &CommentsCtrl{darling.Controller{Request: req, Response: resp, PathParams: []string{topicId}}, AccessController{AccountId: 1}}
	c.Post()
	switch resp.StatusCode {
	case 500:
		t.Error("server error")
	case 201:
		etag, _ := resp.Header["ETag"]
		if etag == "" {
			t.Error("response header has no etag")
		}
		h := md5.New()
		io.WriteString(h, string(resp.Content))
		buffer := bytes.NewBuffer(nil)
		fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
		if etag != buffer.String() {
			t.Error("response's etag is err")
		}
	}
}
