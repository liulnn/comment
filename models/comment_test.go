package models

import (
	"testing"
)

const topicId string = "topic_id"

func addComment() (id int64, err error) {
	source, username, content := "google", "liull", "test content"
	return AddComment(topicId, source, username, content)
}

func deleteComment(ids []int64) {
	for _, v := range ids {
		DeleteComment(topicId, v)
	}
}

func TestAddComment(t *testing.T) {
	id, err := addComment()
	if err != nil {
		t.Fatal(err)
	}
	deleteComment([]int64{id})
}

func TestGetComments(t *testing.T) {
	pageSize, pageNum := 20, 1
	var ids []int64 = make([]int64, pageSize)
	for i := 0; i < pageSize; i++ {
		id, _ := addComment()
		ids = append(ids, id)
	}
	comments, err := GetComments(topicId, pageSize, pageNum)
	if err != nil {
		t.Fatal(err)
	}
	if len(comments) != pageSize {
		t.Fatalf("the comment's length(%d) != pageSize", len(comments))
	}
	deleteComment(ids)
}

func TestGetComment(t *testing.T) {
	commentId, _ := addComment()
	comment, err := GetComment(topicId, commentId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(comment)
	deleteComment([]int64{commentId})
}

func TestDeleteComment(t *testing.T) {
	commentId, _ := addComment()
	success, err := DeleteComment(topicId, commentId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(success)
}
