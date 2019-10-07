package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"github.com/gorilla/mux"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
)

// APIManager : api manager
type APIManager struct {
}

// StartAPI : start the server for API
func StartAPI() {
	log.Get().Println("ApiManager.[StartAPI()] : start api manager")
	router := mux.NewRouter()
	router.HandleFunc(config.Get().Iso.Messaging.Handlers, PostMessageISO).Methods("POST")

	log.Get().Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Get().Iso.Messaging.Port), router))

	log.Get().Println("ApiManager.[StartAPI()] : end api manager")
}

// PostMessageISO : handle post message iso
func PostMessageISO(w http.ResponseWriter, req *http.Request) {
	log.Get().Println("ApiManager.[PostMessageISO(w http.ResponseWriter, req *http.Request)] : start func postmessageiso")
	w.Header().Set("Content-Type", "application/json")
	message, _ := ioutil.ReadAll(req.Body)

	log.Get().Println("ApiManager.[PostMessageISO(w http.ResponseWriter, req *http.Request)] : message received : ", string(message))
	res := Process(message)
	w.Write([]byte(res))
	log.Get().Println("ApiManager.[PostMessageISO(w http.ResponseWriter, req *http.Request)] : end func postmessageiso")
}
