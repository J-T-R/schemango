package main

import (
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

func (ru *Runner) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "up")
}

func (ru *Runner) addSchema(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}
	ru.Schemas["test"] = body
	// ru.updateSubscribers("test")
}

/*
func (ru *Runner) getSchema(w http.ResponseWriter, r *http.Request) {}
*/

/*
func (ru *Runner) updateSubscribers(schema string) error {}
*/

func (ru *Runner) runAPI() {
	r := mux.NewRouter()

	r.HandleFunc("/health", ru.healthCheck).Methods(http.MethodGet)
	r.HandleFunc("/", ru.addSchema).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	server := &http.Server{
		Addr:    "localhost:" + strconv.Itoa(ru.Port),
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
