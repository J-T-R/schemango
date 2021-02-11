package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
	Where are your tests young man?
*/

type Address struct {
	Protocol string
	Hostname string
	Port     int
}

func (a *Address) createPostString() string {
	return a.Protocol + "://" + a.Hostname + ":" + strconv.Itoa(a.Port)
}

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

func main() {
	runner := Runner{}

	// Here we could add a default config getter
	schemas := make(map[string][]byte)

	// Could also we a default config to load
	subscriptions := make(map[string]Address)

	// Port should be from config
	// Although should it just be a full address from config?
	runner.Port = 9000
	runner.Schemas = schemas
	runner.Subscriptions = subscriptions

	runner.runAPI()
}
