package main

import (
	"log"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/processor"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/repository/messages"
)

// Manager : to process the message
type Manager struct{}

// Process : processing json message
func Process(message []byte) (response string) {

	log.Println("Manager[Process(message []byte)] : start processing message in json format")

	// Save the original message from client
	log.Println("Manager[Process(message []byte)] : saving original message")
	messages.Save(message)

	// Change the original partner id and terminal id and save before send to GSP
	log.Println("Manager[Process(message []byte)] : Change the original partner id and terminal id and save before send to GSP")
	clientPartner, clientTerminal := gjson.Get(string(message), "partnerCentralId"), gjson.Get(string(message), "terminalId")
	jsonClientOut, _ := sjson.Set(string(message), "partnerCentralId", config.Get().Gsp.Partner)
	jsonClientOut, _ = sjson.Set(string(jsonClientOut), "terminalId", config.Get().Gsp.Terminal)
	messages.Save([]byte(jsonClientOut))

	// Send message to gsp as ISO message
	log.Println("Manager[Process(message []byte)] : Send message to gsp as ISO message")
	producer := &processor.Message{}
	producer.SetBuilder(&processor.JSONProcessor{})
	result := make(chan []byte)
	go StartDialManager(producer.Process([]byte(jsonClientOut)), result)
	isoMessage := <-result

	// Convert iso message from gsp to json and save the message
	log.Println("Manager[Process(message []byte)] : Convert iso message from gsp to json and save the message")
	producer.SetBuilder(&processor.IsoProcessor{})
	jsonGsp := producer.DecodeMessage(isoMessage)
	messages.Save([]byte(jsonGsp))

	// Change the original partner id and terminal id
	log.Println("Manager[Process(message []byte)] : Change the original partner id and terminal id")
	response, _ = sjson.Set(jsonGsp, "partnerCentralId", clientPartner.String())
	response, _ = sjson.Set(response, "terminalId", clientTerminal.String())
	messages.Save([]byte(response))

	return
}
