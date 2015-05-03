package main

import (
	"comment/controllers"
	"darling"
)

func main() {
	app := darling.NewApp()
	app.Handlers.Add("/comments", &controllers.CommentsCtrl{})
	app.Handlers.Add("/comments/(\\w+)", &controllers.CommentCtrl{})
	app.Run("127.0.0.1", 8001)
}
