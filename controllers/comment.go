package controllers

import (
	"comment/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type CommentEntity struct {
	Content string `form:"content",valid:"Required;Range(1,140)"`
}

type CommentsController struct {
	beego.Controller
}

type CommentCreated struct {
	CommentId int64 `json:"comment_id"`
}

func (c *CommentsController) Post() {
	var ob CommentEntity
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	commentId := models.AddComment(ob)
	c.Data["json"] = &CommentCreated{CommentId: commentId}
	c.ServeJson()
}

type CommentDetail struct {
	Content string
}

type CommentDetailList struct {
	Comments CommentDetail `json:"comments"`
}

func (c *CommentsController) Get() {
	comments := models.GetComments(20, 1)
	c.Data["json"] = &CommentDetailList{Comments: comments}
	c.ServeJson()

}

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Get() {
	comment := models.GetComment(1)
	c.Data["json"] = &CommentDetail{Content: comment.Content}
	c.ServeJson()

}

func (c *CommentController) Delete() {
	err := models.DeleteComment(1)
}
