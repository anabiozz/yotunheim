package handlers

import (
	"github.com/kataras/iris"
)

func HomeHandler(ctx iris.Context) {
	ctx.Redirect("/dashboard")
}
