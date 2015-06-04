package controllers

import (
	"comment/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

type CommentEntity struct {
	Content string `form:"content",valid:"Required;Range(1,140)"`
}

type CommentsController struct {
	beego.Controller
	TopicId string
}

func (c *CommentsController) Prepare() {
	c.TopicId = c.Ctx.Input.Param(":topicId")
}

type CommentCreated struct {
	CommentId int64 `json:"comment_id"`
}

func (c *CommentsController) Post() {
	var ob CommentEntity
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	source := "1"
	username := "1"
	commentId, _ := models.AddComment(c.TopicId, source, username, ob.Content)
	c.Ctx.Output.SetStatus(201)
	c.Ctx.Output.Json(&CommentCreated{CommentId: commentId})
}

func (c *CommentsController) Get() {
	pageSize, pageNum := 20, 1
	comments, _ := models.GetComments(c.TopicId, pageSize, pageNum)
	c.Ctx.Output.Json(comments)
}

type CommentController struct {
	beego.Controller
	TopicId   string
	CommentId int64
}

func (c *CommentController) Prepare() {
	c.TopicId = c.Ctx.Input.Param(":topicId")
	c.CommentId, _ = strconv.ParseInt(c.Ctx.Input.Param(":commentId"), 10, 64)
}

func (c *CommentController) Get() {
	comment, _ := models.GetComment(c.TopicId, c.CommentId)
	c.Data["json"] = comment
	c.ServeJson()

}

func (c *CommentController) Delete() {
	models.DeleteComment(c.TopicId, c.CommentId)
	c.Ctx.Output.SetStatus(204)
}
