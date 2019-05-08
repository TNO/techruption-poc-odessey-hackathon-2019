package mpc

import (
	"encoding/json"
	"log"

	"github.com/odysseyhack/mpan-compute-initiator/nodecomm"
)

type QueryType int

const (
	QUERY_TYPE_INFO QueryType = 0
	QUERY_TYPE_CALC QueryType = 1
)

// Query is the type of query this MPC system supports
type Query struct {
	QueryType       QueryType
	Identifier      int
	Attribute       int
	ClientReference string
}

// Listener loop just does queries sent to its channel
func StartQueryListener() chan Query {
	log.Println("Starting query listener")
	qc := make(chan Query)
	go func() {
		for {
			q := <-qc

			log.Printf("(Query listener) Got query %v", q)
			doQuery(q)
		}
	}()
	return qc
}

// Send a query to the mpc node
func doQuery(query Query) {
	// For now, we just print it (looks nice for the dashboard)
	log.Printf("(doQuery) Received a query with ID %v:\n", query.ClientReference)
	log.Printf("(doQuery)   %v(%v, %v)\n", query.QueryType, query.Identifier, query.Attribute)

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		log.Printf("(doQuery) !!! Error marshaling to string, %v", err)
		return
	}

	log.Printf("(doQuery) Marshalled: %s", jsonQuery)

	nodecomm.Send(append(jsonQuery, '\n'))
}
