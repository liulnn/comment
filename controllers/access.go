package controllers

import (
	"net/http"
)

type AccessController struct {
}

func (c *AccessController) Access(req *http.Request) (accountId int64, err error) {
	username, password, _ := req.BasicAuth()
	if username == "liull" && password == "111111" {
		return 1, nil
	}
	return
}
