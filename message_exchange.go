package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config/db"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/processor"
)

var (
	// MessageClientIn : message from client listen by the server
	MessageClientIn = make(chan []byte)
	// MessageClientOut : response message to client after receipt from ServerDialIn after dial to the iso server
	MessageClientOut = make(chan []byte)
	// ServerDialOut : receipt message from MessageClientIn to prosess for dial
	ServerDialOut = make(chan []byte)
	// ServerDialIn : response message after dial to the iso server
	ServerDialIn = make(chan []byte)
	// JSONMessage : json message
	JSONMessage = make(chan []byte)
	// JSONRelease : json message
	JSONRelease = make(chan []byte)
)

// PostMessageISO : post message iso
func PostMessageISO(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	bytes, _ := ioutil.ReadAll(req.Body)
	// log.Println("received : ", string(b))

	// params := mux.Vars(req)
	producer := &processor.Message{}

	// change the iso bytes to json to save the incoming call from client
	log.Println("MessageExchange : change the iso bytes to json to save the incoming call from client")
	producer.SetBuilder(&processor.IsoProcessor{})
	byteMessage := gjson.Get(string(bytes), "message")
	jsonClient := producer.DecodeMessage([]byte(byteMessage.String()))
	db.Save([]byte(jsonClient))
	clientPartner, clientTerminal := gjson.Get(jsonClient, "partnerCentralId"), gjson.Get(jsonClient, "terminalId")

	// send the iso message after change the field from partner to gsp
	log.Println("MessageExchange : send the iso message after change the field from partner to gsp")
	log.Println("MessageExchange[PostMessageISO] json client : ", jsonClient)
	repJSONClient, _ := sjson.Set(jsonClient, "partnerCentralId", config.Get().Gsp.Partner)
	repJSONClient, _ = sjson.Set(repJSONClient, "terminalId", config.Get().Gsp.Terminal)
	log.Println("MessageExchange[PostMessageISO] json client replace value : ", repJSONClient)
	db.Save([]byte(repJSONClient))

	producer.SetBuilder(&processor.JSONProcessor{})
	gspIsoBytes := producer.Process([]byte(repJSONClient))
	MessageClientIn <- gspIsoBytes

	// receive the message from gsp
	log.Println("MessageExchange : receive the message from gsp")
	message := <-MessageClientOut
	producer.SetBuilder(&processor.IsoProcessor{})
	jsonGsp := producer.DecodeMessage(message)
	db.Save([]byte(jsonGsp))

	// send the iso message after change the field from gsp to partner
	log.Println("MessageExchange : send the iso message after change the field from gsp to partner")
	repJSONGsp, _ := sjson.Set(jsonGsp, "partnerCentralId", clientPartner.String())
	repJSONGsp, _ = sjson.Set(repJSONGsp, "terminalId", clientTerminal.String())
	db.Save([]byte(repJSONGsp))
	producer.SetBuilder(&processor.JSONProcessor{})
	partnerIsoBytes := producer.Process([]byte(repJSONGsp))
	log.Println("MessageExchange data to send : ", string(partnerIsoBytes))
	jsonData := map[string]string{"message": string(partnerIsoBytes)}
	jsonValue, _ := json.Marshal(jsonData)
	w.Write(jsonValue)
}
