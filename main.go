package main

// var wg = sync.WaitGroup{}

// GetMessageISO : get message iso
// func GetMessageISO(w http.ResponseWriter, req *http.Request) {

// 	params := mux.Vars(req)
// 	producer := &processor.Message{}
// 	jsonP := &processor.JSONProcessor{}
// 	producer.SetBuilder(jsonP)
// 	MessageClientIn <- producer.Process([]byte(params["iso"]))
// 	// connection.Write([]byte(strings.TrimRight(message, "\n")))
// 	message := <-ServerDialIn
// 	iso := &processor.IsoProcessor{}
// 	producer.SetBuilder(iso)
// 	fmt.Println(producer.DecodeMessage(message))
// 	json.NewEncoder(w).Encode(producer.DecodeMessage(message))
// }

func main() {

	StartAPI()
	// log.Println("starting biller payment ...")

	// wg.Add(3)
	// go StartListenerServer()
	// go StartDialManager()

	// go func() {
	// 	log.Println("[MessageExchange] : start the function")
	// 	defer wg.Done()
	// 	for {
	// 		select {
	// 		case message := <-MessageClientIn:
	// 			log.Println("[MessageExchange] : received message from client ")
	// 			ServerDialOut <- message
	// 		case message := <-ServerDialIn:
	// 			producer := &processor.Message{}
	// 			iso := &processor.IsoProcessor{}
	// 			producer.SetBuilder(iso)
	// 			jsonResult := producer.DecodeMessage(message)
	// 			db.Save([]byte(jsonResult))
	// 			fmt.Println("result : ", jsonResult)
	// 		}
	// 	}
	// }()

	// go func() {
	// 	log.Println("[MessageExchange] : start the function")
	// 	defer wg.Done()
	// 	for {
	// 		select {
	// 		case message := <-MessageClientIn:
	// 			log.Println("[MessageExchange] : received message from client ")
	// 			ServerDialOut <- message
	// 		case message := <-ServerDialIn:
	// 			MessageClientOut <- message
	// 		}
	// 	}
	// }()

	// router := mux.NewRouter()
	// router.HandleFunc(config.Get().Iso.Messaging.Handlers, PostMessageISO).Methods("POST")

	// log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Get().Iso.Messaging.IP,
	// 	config.Get().Iso.Messaging.Port), router))

	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	message, _ := reader.ReadString('\n')
	// 	producer := &processor.Message{}
	// 	json := &processor.JSONProcessor{}
	// 	db.Save([]byte(strings.TrimRight(message, "\n")))
	// 	producer.SetBuilder(json)
	// 	// MessageClientIn <- producer.Process([]byte(message))
	// 	// connection.Write([]byte(strings.TrimRight(message, "\n")))
	// 	MessageClientIn <- producer.Process([]byte(strings.TrimRight(message, "\n")))
	// }

	// wg.Wait()

	// fmt.Println("starting")
	// ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// collection := config.GetDB().Collection("messages")
	// res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// collection.InsertOne(ctx, "arip")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// id := res.InsertedID
	// fmt.Println(id)

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.Get().Db.URI))

	// collection := client.Database(config.Get().Db.Document).Collection("messages")
	// ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	// res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// id := res.InsertedID
	// fmt.Println(id)

	// collection := config.GetDB().Collection("messages")
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// id := res.InsertedID
	// fmt.Println(id)

	// wg.Wait()
}
