package main

import (
	"encoding/json"
	"fmt"
	"haki/common"
	"haki/config"
	"haki/websocket"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mux "github.com/julienschmidt/httprouter"
)

var connection *websocket.Connection

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  mux.Handle
}

type Routes []Route

// Index xx
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// SubsIndex xx
func SubsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	subs := connection.Subs

	fmt.Println(subs)
	if err := json.NewEncoder(w).Encode(subs); err != nil {
		panic(err)
	}
}

// SubsCreate xx
func SubsCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	var sub_key common.Sub
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &sub_key); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(sub_key.Sub)

	// s2 = append(s2, 5)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(sub_key); err != nil {
		panic(err)
	}
}

// SubsDelete 取消订阅
func SubsDelete(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	var sub_key common.Sub
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &sub_key); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	connection.Subs = common.DeleteByName(connection.Subs, sub_key.Sub)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(sub_key); err != nil {
		panic(err)
	}

}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"SubsIndex",
		"GET",
		"/subs",
		SubsIndex,
	},
	Route{
		"SubsCreate",
		"POST",
		"/subs",
		SubsCreate,
	},
	// Route{
	// 	"TodoDownload",
	// 	"GET",
	// 	"/todos.json",
	// 	TodoDownload,
	// },
}

func NewRouter() *mux.Router {

	router := mux.New()

	for _, route := range routes {

		/*
			var handler mux.Handle

			handler = Logger(router, route.Method, route.Pattern, route.Name)
		*/

		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}

func main() {

	// client := amqp.NewClient(config.RABBITMQ_URL, config.RABBITMQ_EXCHANGE)

	// defer client.Close()

	// mongoClient := mongodb.NewDBClient(config.MONGODB_URL, config.MONGODB_DATABASE)
	// database := mongoClient.GetDatabase()
	// collection := database.Collection("collection")

	// defer mongoClient.CloseDatabase()

	connection = websocket.NewConnection(config.WS_COIN_URL)
	connection.Connect()
	connection.HeartBeat()

	defer connection.CloseConnection()

	connection.Subscribe(map[string]string{
		"sub": "market.htusdt.kline.1min",
		"id":  "1"})

	// connection.Subscribe(map[string]string{
	// 	"sub": "market.btcusdt.kline.1min",
	// 	"id":  "1"})

	// connection.Watch(func(msg []byte) {
	// 	gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
	// 	data, _ := ioutil.ReadAll(gzipreader)
	// 	var resp map[string]interface{}
	// 	json.Unmarshal(data, &resp)

	// 	if resp["ping"] != nil {
	// 		connection.Ws.WriteJSON(map[string]interface{}{"pong": resp["ping"]})
	// 	} else if resp["ch"] != nil {
	// 		// kv := strings.Split(resp["ch"].(string), ".")
	// 		// fmt.Println(kv)
	// 		var ticker common.Ticker
	// 		json.Unmarshal(data, &ticker)
	// 		fmt.Println(ticker)
	// 		// _, _ = collection.InsertOne(context.TODO(), ticker)
	// 	} else {
	// 	}
	// })

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
