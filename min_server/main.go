package main

import (
	"min"
	"net/http"
)

func main() {
	engine := min.New()

	group := engine.Group("/api")

	group.Use(min.Recovery(), min.Logger())

	group.GET("/echo/*name", func(ctx *min.Context) {
		ctx.JSON(http.StatusOK, min.H{"echo": ctx.Params["name"]})
	})
	group.POST("/ping", func(ctx *min.Context) {
		ctx.JSON(http.StatusOK, "pone")
	})

	group.POST("/panic", func(ctx *min.Context) {
		panic("panic")
	})

	engine.Run(":9999")
}
