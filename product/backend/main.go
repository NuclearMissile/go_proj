package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	tmplate := iris.
		HTML("backend/web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmplate)
	app.StaticWeb("/assets", "backend/web/assets")
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().
			GetStringDefault("message", "Error occurred."))
		ctx.ViewLayout("")
		ctx.View("share/error.html")
	})
	app.Run(iris.Addr("localhost:8080"),
		iris.WithOptimizations,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
