package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/odysseyhack/mpan-compute-initiator/mpc"
	"github.com/odysseyhack/techruption-multi-party-all-night/phone-interface/smartcontract"
)

var ch chan mpc.Query

func main() {
	results = make(map[string]*Result)

	ch = smartcontract.WaitForQueries()

	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/getresult", getResultHandler)
	http.HandleFunc("/setresult", setResultHandler)

	log.Println("Setup complete, listening...")

	log.Fatal(http.ListenAndServe(":80", nil))
}

var results map[string]*Result

type Result []byte

func queryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a query")

	// Android client is interesting: it sends json. Parse it
	result := struct{ Function, Identifier, Attribute string }{}
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		badRequest(w, fmt.Sprint(err))
		return
	}
	identifier, _ := strconv.Atoi(result.Identifier)
	attribute, _ := strconv.Atoi(result.Attribute)

	switch result.Function {
	case "info":
		queryHandlerInfo(w, identifier)

	case "calc":
		queryHandlerCalc(w, identifier, attribute)

	default:
		badRequest(w, "Bad request: function must be info or calc")
		return
	}
}

func queryHandlerInfo(w http.ResponseWriter, identifier int) {
	id := uuid.New().String()
	q := &mpc.Query{
		QueryType:       mpc.QUERY_TYPE_INFO,
		Identifier:      identifier,
		Attribute:       0,
		ClientReference: id,
	}
	blockchainSubmit(q)
	log.Println("INFO created")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"id":"%v"}`, id)
}

func queryHandlerCalc(w http.ResponseWriter, identifier int, attribute int) {
	id := uuid.New().String()
	q := &mpc.Query{
		QueryType:       mpc.QUERY_TYPE_CALC,
		Identifier:      identifier,
		Attribute:       attribute,
		ClientReference: id,
	}
	blockchainSubmit(q)
	log.Println("CALC created")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"id":"%v"}`, id)
}

func blockchainSubmit(query *mpc.Query) {
	log.Printf("Submitting %+v", *query)
	ch <- *query
}

func getResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a request for result")

	r.ParseForm()
	if len(r.Form["id"]) == 0 {
		badRequest(w, "Bad request: parameter id missing")
		return
	}

	log.Printf("Will wait for result %v", r.Form["id"][0])

	for {
		result := results[strings.TrimSpace(r.Form["id"][0])]
		if result != nil {
			log.Printf("Content: %s", *result)
			w.Write(append([]byte(*result), '\n'))
			log.Printf("Gave result %v", r.Form["id"][0])
			return
		}
		select {
		case <-time.After(time.Second):
		case <-r.Context().Done():
			log.Printf("Client gave up on %v", r.Form["id"][0])
			return
		}
	}
}

func setResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a submission of a result")

	r.ParseForm()
	if len(r.PostForm["id"]) == 0 || len(r.PostForm["result"]) == 0 {
		badRequest(w, "Bad request: parameter id or result missing")
		return
	}

	id := strings.TrimSpace(r.PostForm["id"][0])
	result := r.PostForm["result"][0]

	if len(id) == 0 || len(result) == 0 {
		badRequest(w, "Bad request: parameter id == \"\" or result == \"\"")
		return
	}

	rs := Result(result)
	results[id] = &rs

	log.Printf("Created %v -> %v", id, rs)
	w.WriteHeader(http.StatusCreated)
}

func badRequest(w http.ResponseWriter, message string) {
	log.Println("Bad request: " + message)
	http.Error(w, "Bad request: "+message, http.StatusBadRequest)
}
