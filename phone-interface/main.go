package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/odysseyhack/mpan-compute-initiator/mpc"
)

func main() {
	results = make(map[uint64]*Result)

	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/getresult", getResultHandler)
	http.HandleFunc("/setresult", setResultHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}

var results map[uint64]*Result

type Result []byte

func queryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.PostForm["function"]) == 0 || len(r.PostForm["identifier"]) == 0 {
		http.Error(w, "Bad request: parameter function or identifier missing", http.StatusBadRequest)
	}

	switch r.PostForm["function"][0] {
	case "info":
		id := time.Now().Unix()
		identifier, _ := strconv.Atoi(r.PostForm["identifier"][0])
		q := &mpc.Query{
			QueryType:  mpc.QUERY_TYPE_INFO,
			Identifier: identifier,
			Attribute:  0,
			QueryId:    big.NewInt(id),
		}
		blockchainSubmit(q)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "id=%v\n", id)

	case "mpc": // not called yet?
		if len(r.PostForm["attribute"]) == 0 {
			http.Error(w, "Bad request: parameter attribute missing", http.StatusBadRequest)
		}
		identifier, _ := strconv.Atoi(r.PostForm["identifier"][0])
		attribute, _ := strconv.Atoi(r.PostForm["attribute"][0])

		id := time.Now().Unix()
		q := &mpc.Query{
			QueryType:  mpc.QUERY_TYPE_CALC,
			Identifier: identifier,
			Attribute:  attribute,
			QueryId:    big.NewInt(id),
		}
		blockchainSubmit(q)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "id=%v\n", id)

	default:
		http.Error(w, "Bad request: function must be info or mpc", http.StatusBadRequest)
	}
}

func blockchainSubmit(query *mpc.Query) {

}

func getResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["id"]) == 0 {
		http.Error(w, "Bad request: parameter id missing", http.StatusBadRequest)
	}

	id, _ := strconv.ParseUint(r.PostForm["id"][0], 10, 64)
	if result := results[id]; result == nil {
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

	id, _ := strconv.ParseUint(r.PostForm["id"][0], 10, 64)
	result := r.PostForm["result"][0]

	if id == 0 || len(result) == 0 {
		http.Error(w, "Bad request: parameter id == 0 or result == \"\"", http.StatusBadRequest)
	}

	rs := Result(result)
	results[id] = &rs

	w.WriteHeader(http.StatusCreated)
}
