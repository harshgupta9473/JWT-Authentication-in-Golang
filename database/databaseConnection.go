package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	err:= godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}

	MongoDb:= os.Getenv("MONGODB_URL")
	ctx ,cancel:= context.WithTimeout(context.Background(),10*time.Second)
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(MongoDb))
	if err!=nil{
		log.Fatal(err)
	}
	defer cancel()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client =DBinstance()

func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	 collection:= client.Database("JWTHehe").Collection(collectionName)
	 return collection
}