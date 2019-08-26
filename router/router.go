package router

import (
	"encoding/json"
	"fmt"
	"haki/common"
	"haki/websocket"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//	"net/http"

//Route 路由结构
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

// Routes 路由数组
type Routes []Route

// NewRouter xx
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes {

		/*
			var handler mux.Handle

			handler = Logger(router, route.Method, route.Pattern, route.Name)
		*/
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}

// Index xx
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// SubsIndex xx
func SubsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	subs := websocket.WsConnection.Subs

	fmt.Println(subs)
	if err := json.NewEncoder(w).Encode(subs); err != nil {
		panic(err)
	}
}

// SubsCreate xx
func SubsCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var subKey common.Sub
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &subKey); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(subKey.Sub)

	// s2 = append(s2, 5)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(subKey); err != nil {
		panic(err)
	}
}

// SubsDelete 取消订阅
func SubsDelete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var subKey common.Sub
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &subKey); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	websocket.WsConnection.Subs = common.DeleteByName(websocket.WsConnection.Subs, subKey.Sub)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(subKey); err != nil {
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
