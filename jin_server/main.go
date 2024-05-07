package main

import (
	"jin"
	"net/http"
)

func main() {
	engine := jin.New()

	group := engine.Group("/api")

	group.Use(jin.Recovery(), jin.Logger())

	group.GET("/echo/*name", func(ctx *jin.Context) {
		ctx.JSON(http.StatusOK, jin.H{"echo": ctx.Params["name"]})
	})
	group.POST("/ping", func(ctx *jin.Context) {
		ctx.JSON(http.StatusOK, "pone")
	})

	group.POST("/panic", func(ctx *jin.Context) {
		panic("panic")
	})

	engine.Run(":9999")
}
