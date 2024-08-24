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

	clientOptions := options.Client().ApplyURI("mongodb+srv://nnm23mc060:le3uh4DDWIjuDymr@cluster0.ghe7tvk.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

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

	log.Println("Connected to MongoDB Atlas!")

	controllers.InitProductController(client, "restapiforgo", "products")

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
