package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/kataras/iris"
)

type User struct {
	Name     string  `json:"name"`
	BMI      float32 `json:"BMI"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	City     string  `json:"city"`
	Married  bool    `json:"married"`
	Index    int     `json:"index"`
}

func main() {
	app := iris.New()
	tmpl := iris.HTML("./public", ".html")
	tmpl.Layout("index.html")
	app.RegisterView(tmpl)

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.Redirect("/dashboard")
	})

	app.Handle("GET", "/dashboard", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	app.Handle("GET", "/api/get-json", func(ctx iris.Context) {
		type data []interface{}
		users, _ := ioutil.ReadFile("./public/user.json")
		dec := json.NewDecoder(bytes.NewReader(users))
		var d data
		dec.Decode(&d)
		if _, err := ctx.JSON(&d); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	app.StaticWeb("/", "./public")
	app.Run(
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
