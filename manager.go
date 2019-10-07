package main

import (
	"bytes"
	"time"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/repository/messages"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/processor"
)

// Manager : to process the message
type Manager struct{}

// Process : processing json message
func Process(message []byte) (response string) {

	log.Get().Println("Manager[Process(message []byte)] : start processing message in json format")

	producer := &processor.Message{}

	// Save the original message from client
	log.Get().Println("Manager[Process(message []byte)] : saving original message from client to 2Pay")
	jsonClientIn, _ := sjson.Set(string(message), "createdDate", time.Now())
	messages.Save([]byte(jsonClientIn))

	// Change the original partner id and terminal id and save before send to GSP
	log.Get().Println("Manager[Process(message []byte)] : Change the original partner id and terminal id and save before send to GSP")
	clientPartner, clientTerminal, bankcode := gjson.Get(string(message), "partnerCentralId"), gjson.Get(string(message), "terminalId"), gjson.Get(string(message), "bankCode")
	var unique bytes.Buffer
	unique.WriteString(clientPartner.String())
	unique.WriteString(bankcode.String())
	log.Get().Println("unique value : ", unique.String())
	if unique.String() != "09000010850001" {
		log.Get().Println("return because partner central id and bank code is not same")
		return
	}

	json2PayOut, _ := sjson.Set(string(message), "partnerCentralId", config.Get().Gsp.Partner)
	json2PayOut, _ = sjson.Set(string(json2PayOut), "terminalId", config.Get().Gsp.Terminal)
	json2PayOut, _ = sjson.Set(string(json2PayOut), "bankCode", config.Get().Gsp.Bankcode)
	json2PayOut, _ = sjson.Set(string(json2PayOut), "createdDate", time.Now())

	producer.SetBuilder(&processor.JSONProcessor{})
	result := make(chan []byte)
	iso2PayToGSP := producer.Process([]byte(json2PayOut))
	json2PayOut, _ = sjson.Set(string(json2PayOut), "payload", string(iso2PayToGSP))
	messages.Save([]byte(json2PayOut))

	// Send message to gsp as ISO message
	log.Get().Println("Manager[Process(message []byte)] : Send message to gsp as ISO message")
	go StartDialManager(iso2PayToGSP, result)
	isoGSPTo2Pay := <-result

	// Convert iso message from gsp to json and save the message
	log.Get().Println("Manager[Process(message []byte)] : Convert iso message from gsp to json and save the message")
	producer.SetBuilder(&processor.IsoProcessor{})
	jsonGsp := producer.DecodeMessage(isoGSPTo2Pay)
	jsonGspRev, _ := sjson.Set(string(jsonGsp), "createdDate", time.Now())
	jsonGspRev, _ = sjson.Set(string(jsonGspRev), "payload", string(isoGSPTo2Pay))
	messages.Save([]byte(jsonGspRev))

	// Change the original partner id and terminal id
	log.Get().Println("Manager[Process(message []byte)] : Change the original partner id and terminal id")
	response, _ = sjson.Set(jsonGsp, "partnerCentralId", clientPartner.String())
	response, _ = sjson.Set(response, "terminalId", clientTerminal.String())
	response, _ = sjson.Set(response, "bankCode", bankcode.String())
	jsonClientOut, _ := sjson.Set(string(response), "createdDate", time.Now())
	producer.SetBuilder(&processor.JSONProcessor{})
	iso2PayToClient := producer.Process([]byte(jsonClientOut))
	jsonClientOut, _ = sjson.Set(string(jsonClientOut), "payload", string(iso2PayToClient))
	messages.Save([]byte(jsonClientOut))

	return response
}
