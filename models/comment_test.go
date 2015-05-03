package models

import (
	"log"
	"testing"
)

func TestAddComment(t *testing.T) {
	content := "this is content"
	comment, err := AddComment("1", 1, content)
	if err != nil {
		t.Error(err)
	}
	log.Println(comment)
}

func TestGetComments(t *testing.T) {
	comments, err := GetComments("1", 10, 0)
	if err != nil {
		t.Error(err)
	}
	log.Println(comments)
}
func TestGetComment(t *testing.T) {
	comment, err := GetComment("1", 1)
	if err != nil {
		t.Error(err)
	}
	log.Println(comment)
}
func TestDeleteComment(t *testing.T) {
	err := DeleteComment("1", 1)
	if err != nil {
		t.Error(err)
	}
}
