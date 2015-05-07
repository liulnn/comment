package controllers

import (
	"net/http"
	"testing"
)

func TestAccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	req.SetBasicAuth("liull", "111111")
	c := new(AccessController)
	err := c.Access(req)
	if err != nil {
		t.Error(err)
	}
	if c.AccountId == 0 {
		t.Error("account not found")
	}
}
