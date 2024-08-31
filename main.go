package main

import (
	"context"
	"log"
	"restapiingo/controllers"
	"restapiingo/routes"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("your_mongo_uri")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	controllers.InitProductController(client, "restapiforgo", "products")
	controllers.InitAuthController(client, "restapiforgo", "users")

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
