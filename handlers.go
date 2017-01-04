package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func BaseHandeler(w http.ResponseWriter, r *http.Request, t interface{}) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal(body, t); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New qos"}' http://localhost:8080/qos/client

*/
func QosClient(w http.ResponseWriter, r *http.Request) {
	var qosClient QosClientT
	// var body []byte

	// BaseHandeler(w, r, &qosClient)

	// t := RepoCreateTodo(qosClient)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	// }

	if err := json.NewDecoder(r.Body).Decode(&qosClient); err != nil {
		log.Println("Could not decode body:", err)
	}

	t := RepoCreateTodo(qosClient)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

}
