package main

import (
	"github.com/savsgio/atreugo/v10"
)

func main() {
	config := &atreugo.Config{Addr: "0.0.0.0:8000"}
	server := atreugo.New(config)

	api := server.NewGroupPath("/api")
	api.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("API")
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
