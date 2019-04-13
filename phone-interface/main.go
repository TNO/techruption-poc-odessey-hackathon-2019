package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	r.ParseForm()
	if len(r.PostForm["function"]) == 0 || len(r.PostForm["identifier"]) == 0 {
		log.Println("Bad request: parameter function or identifier missing")
		http.Error(w, "Bad request: parameter function or identifier missing", http.StatusBadRequest)
		return
	}

	switch r.PostForm["function"][0] {
	case "info":
		identifier, _ := strconv.Atoi(r.PostForm["identifier"][0])
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
		fmt.Fprintf(w, "%v\n", id)

	case "calc": // not called yet?
		if len(r.PostForm["attribute"]) == 0 {
			log.Println("Bad request: parameter attribute missing")
			http.Error(w, "Bad request: parameter attribute missing", http.StatusBadRequest)
			return
		}
		identifier, _ := strconv.Atoi(r.PostForm["identifier"][0])
		attribute, _ := strconv.Atoi(r.PostForm["attribute"][0])

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
		fmt.Fprintf(w, "%v\n", id)

	default:
		log.Println("Bad request: function must be info or calc")
		http.Error(w, "Bad request: function must be info or calc", http.StatusBadRequest)
		return
	}
}

func blockchainSubmit(query *mpc.Query) {
	log.Printf("Submitting %+v", *query)
	ch <- *query
}

func getResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a request for result")

	r.ParseForm()
	if len(r.Form["id"]) == 0 {
		log.Println("Bad request: parameter id missing")
		http.Error(w, "Bad request: parameter id missing", http.StatusBadRequest)
		return
	}

	if result := results[strings.TrimSpace(r.Form["id"][0])]; result == nil {
		log.Println("No content")
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Printf("Content: %s", *result)
		w.Write(append([]byte(*result), '\n'))
	}
}

func setResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a submission of a result")

	r.ParseForm()
	if len(r.PostForm["id"]) == 0 || len(r.PostForm["result"]) == 0 {
		log.Println("Bad request: parameter id or result missing")
		http.Error(w, "Bad request: parameter id or result missing", http.StatusBadRequest)
		return
	}

	id := strings.TrimSpace(r.PostForm["id"][0])
	result := r.PostForm["result"][0]

	if len(id) == 0 || len(result) == 0 {
		log.Println("Bad request: parameter id == 0 or result == \"\"")
		http.Error(w, "Bad request: parameter id == 0 or result == \"\"", http.StatusBadRequest)
		return
	}

	rs := Result(result)
	results[id] = &rs

	log.Printf("Created %v -> %v", id, rs)
	w.WriteHeader(http.StatusCreated)
}
