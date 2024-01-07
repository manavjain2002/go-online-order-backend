package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database = connectToDatabase()

func connectToDatabase() *mongo.Database {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
	   }
	uri := os.Getenv("URI")
	dbname := "onlineApp"

	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), options)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Mongo Connection Success")

	db := client.Database(dbname)
	return db
}
