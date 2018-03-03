package handlers

import (
	"github.com/kataras/iris"
)

func DashboardHandler(ctx iris.Context) {
	if err := ctx.View("index.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}
