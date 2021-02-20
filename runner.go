package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Runner struct {
	Schemas       map[string][]byte
	Subscriptions map[string]Address
	Port          int
}

type ExpectedBody struct {
	Key    string `json:key`
	Schema []byte `json:schema`
}

func (ru *Runner) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "up")
}

func (ru *Runner) addSchema(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)

	expectedBody := ExpectedBody{}
	err := json.Unmarshal(body, expectedBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ru.Schemas[expectedBody.Key] = expectedBody.Schema
	ru.updateSubscribers(expectedBody.Key)

	w.WriteHeader(http.StatusOK)
}

func (ru *Runner) getSchema(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	schema := ru.Schemas[id]
	if len(schema) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Write(ru.Schemas[id])
		w.WriteHeader(http.StatusOK)
	}
}

/*
func (ru *Runner) updateSubscribers(schema string) error {}
*/

func (ru *Runner) runAPI() {
	r := mux.NewRouter()

	r.HandleFunc("/health", ru.healthCheck).Methods(http.MethodGet)
	r.HandleFunc("/schema", ru.addSchema).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	r.HandleFunc("/schema/{id}", ru.getSchema).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    "localhost:" + strconv.Itoa(ru.Port),
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
