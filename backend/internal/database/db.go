package database

import (
	"context"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ultimateshadsform/pi-dashboard/internal/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var cli *mongo.Client

func Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(helpers.GetEnv("MONGO_URI", "mongodb://localhost:27017")))
	if err != nil {
		return err
	}

	cli = client
	db = client.Database("pi-dashboard")
	return nil
}

func GetDB() *mongo.Database {
	if db == nil {
		panic("Database not initialized")
	}
	return db
}

func GetClient() *mongo.Client {
	if cli == nil {
		panic("Database not initialized")
	}
	return cli
}

func CloseDB() {
	if cli != nil {
		if err := cli.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}
}
