package demo

import (
	"fmt"
	"net/http"
)

func demo() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer return hello world
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
