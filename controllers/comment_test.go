package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type CommentInfo struct {
	content string
}

func TestCommentsCtrlPost(t *testing.T) {
	client := new(http.Client)
	var topicId string = "1"
	path := fmt.Sprintf("/topics/%s/comments", topicId)
	data, _ := json.Marshal(&CommentInfo{content: "test"})
	resp, err := client.Post(path, "application/json", data)
	if resp.StatusCode != 201 {
		t.Error(err)
	}
}

func TestCommentsCtrlGet(t *testing.T) {
}

func TestCommentCtrlGet(t *testing.T) {

}
func TestCommentCtrlDelete(t *testing.T) {

}
