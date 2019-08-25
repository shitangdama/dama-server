package main

import (
	"encoding/json"
	"fmt"
	"haki/common"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//	"net/http"

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  httprouter.Handle
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
func SubsCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
func SubsDelete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
