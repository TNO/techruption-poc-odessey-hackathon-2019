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

	log.Fatal(http.ListenAndServe(":80", nil))
}

var results map[string]*Result

type Result []byte

func queryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.PostForm["function"]) == 0 || len(r.PostForm["identifier"]) == 0 {
		http.Error(w, "Bad request: parameter function or identifier missing", http.StatusBadRequest)
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
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "id=%v\n", id)

	case "calc": // not called yet?
		if len(r.PostForm["attribute"]) == 0 {
			http.Error(w, "Bad request: parameter attribute missing", http.StatusBadRequest)
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
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%v\n", id)

	default:
		http.Error(w, "Bad request: function must be info or mpc", http.StatusBadRequest)
	}
}

func blockchainSubmit(query *mpc.Query) {
	ch <- *query
}

func getResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["id"]) == 0 {
		http.Error(w, "Bad request: parameter id missing", http.StatusBadRequest)
	}

	if result := results[strings.TrimSpace(r.PostForm["id"][0])]; result == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Write([]byte(*result))
	}
}

func setResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.PostForm["id"]) == 0 || len(r.PostForm["result"]) == 0 {
		http.Error(w, "Bad request: parameter id or result missing", http.StatusBadRequest)
	}

	id := strings.TrimSpace(r.PostForm["id"][0])
	result := r.PostForm["result"][0]

	if len(id) == 0 || len(result) == 0 {
		http.Error(w, "Bad request: parameter id == 0 or result == \"\"", http.StatusBadRequest)
	}

	rs := Result(result)
	results[id] = &rs

	w.WriteHeader(http.StatusCreated)
}
