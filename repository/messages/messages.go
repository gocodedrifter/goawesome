package messages

import (
	"context"
	"strings"
	"time"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"

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
		log.Get().Println("messages.Save(json []byte) : unable to insert data : ", err.Error())
	}
	log.Get().Println(id.InsertedID)
}

// FindData : find data using subscriber id, primary account number, and mti
func FindData(mti string, subscriberID string, primaryAccountNumber string) (message *basic.Message, err error) {
	collection := db.GetDB().Collection("messages_simulator")
	subscriberID = strings.Replace(subscriberID, " ", "", -1)
	log.Get().Println(mti, subscriberID, primaryAccountNumber)
	filter := bson.M{"mti": mti, "subscriberId": subscriberID, "primaryAccountNumber": primaryAccountNumber}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err = collection.FindOne(ctx, filter).Decode(&message); err != nil {
		log.Get().Println("error : ", err.Error())
	}
	return
}
