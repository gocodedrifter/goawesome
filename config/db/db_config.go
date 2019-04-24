package db

import (
	"context"
	"log"
	"sync"
	"time"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config : db config
type Config struct {
	*mongo.Database
}

var dbConfig *Config

var syncOnce sync.Once

// GetDB : get db
func GetDB() *Config {
	syncOnce.Do(func() {
		dbConfig = loadDBConfig()
	})
	return dbConfig
}

// loadDBConfig : load db config
func loadDBConfig() *Config {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Get().Db.URI))
	if err != nil {
		panic(err.Error())
	}

	billerSystem := client.Database(config.Get().Db.Document)
	return &Config{billerSystem}
}

// Save : save json
func Save(json []byte) {
	var dataToSave interface{}
	bson.UnmarshalExtJSON(json, true, &dataToSave)
	collection := GetDB().Collection("messages")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id, err := collection.InsertOne(ctx, &dataToSave)
	if err != nil {
		log.Println("unable to insert data : ", err.Error())
	}
	log.Println(id.InsertedID)
}
