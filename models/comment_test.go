package models

import (
	"testing"
)

const topicId string = "topic_id"

func TestAddComment(t *testing.T) {
	source, username, content := "google", "liull", "test content"
	id, err := AddComment(topicId, source, username, content)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestGetComments(t *testing.T) {
	pageSize, pageNum := 20, 1
	comments, err := GetComments(topicId, pageSize, pageNum)
	if err != nil {
		t.Fatal(err)
	}
	if len(comments) > pageSize {
		t.Fatal("the comment's length > pageSize")
	}
}

func TestGetComment(t *testing.T) {
	var commentId int64 = 1
	comment, err := GetComment(topicId, commentId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(comment)
}

func TestDeleteComment(t *testing.T) {
	var commentId int64 = 1
	success, err := DeleteComment(topicId, commentId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(success)
}
