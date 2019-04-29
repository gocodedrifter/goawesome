package messages

import (
	"context"
	"log"
	"time"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Save : save json
func Save(json []byte) {
	var dataToSave interface{}
	bson.UnmarshalExtJSON(json, true, &dataToSave)
	collection := db.GetDB().Collection("messages")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id, err := collection.InsertOne(ctx, &dataToSave)
	if err != nil {
		log.Println("messages.Save(json []byte) : unable to insert data : ", err.Error())
	}
	log.Println(id.InsertedID)
}
