package controllers

import (
	"net/http"
)

type AccessController struct {
	AccountId int64
}

func (c *AccessController) Access(req *http.Request) (err error) {
	username, password, _ := req.BasicAuth()
	if username == "liull" && password == "111111" {
		c.AccountId = 1
		return nil
	}
	return nil
}
