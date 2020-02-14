package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mikeStr8s/delvr-api/db"
	"github.com/savsgio/atreugo/v10"
	"go.mongodb.org/mongo-driver/bson"
)

type Ability struct {
	Name string
}

func main() {
	config := &atreugo.Config{Addr: "0.0.0.0:8000"}
	server := atreugo.New(config)

	api := server.NewGroupPath("/api")
	api.GET("/", func(ctx *atreugo.RequestCtx) error {
		database := db.ConnectDatabase()
		abilities := database.Collection("ability")
		filter := bson.D{{"name", "Strength"}}

		var result Ability
		err := abilities.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.TextResponse(fmt.Sprintf("Found Single Document: %+v\n", result))
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
