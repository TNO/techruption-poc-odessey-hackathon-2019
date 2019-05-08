package smartcontract

//go:generate abigen -abi governance.abi -pkg smartcontract -type Governance -out governance.go
//go:generate abigen -abi gatekeeper.abi -pkg smartcontract -type Gatekeeper -out gatekeeper.go

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/odysseyhack/mpan-compute-initiator/mpc"
)

const (
	SMARTCONTRACT_ADDRESS = "0xa652605f3794d5cd868aa5f295e60fae924fe836"
	ETHEREUM_URL          = "ws://127.0.0.1:8546"
)

func WaitForQuery(queryChan chan mpc.Query) {
	log.Println("(SmartContract) Starting WaitForQuery")

	conn, err := ethclient.Dial(ETHEREUM_URL)
	if err != nil {
		log.Fatalf("(SmartContract) Failed to connect to the Ethereum client: %v", err)
	}

	governance, err := NewGovernance(common.HexToAddress(SMARTCONTRACT_ADDRESS), conn)
	if err != nil {
		log.Fatalf("(SmartContract) Failed to instantiate the Governance contract: %v", err)
	}

	computeAddress, err := governance.GetGatekeeperAddress(nil)
	if err != nil {
		log.Fatalf("(SmartContract) Failed to talk to the Governance contract: %v", err)
	}

	// Instantiate the contract and display its name
	gatekeeper, err := NewGatekeeper(computeAddress, conn)
	if err != nil {
		log.Fatalf("(SmartContract) Failed to instantiate the Gatekeeper contract: %v", err)
	}

	log.Println("(SmartContract) Contracts found")

	var realEventChannel = make(chan *GatekeeperQuery)
	var blockNumber uint64 = 1
	opts := &bind.WatchOpts{}
	opts.Start = &blockNumber
	_, err = gatekeeper.WatchQuery(opts, realEventChannel, nil)
	if err != nil {
		log.Fatal("(SmartContract) Unable to subscribe to event!", err)
	}

	log.Println("(SmartContract) Waiting for smart-contract approved queries.")

	for {

		event := <-realEventChannel
		log.Println("(SmartContract) Received a real query request from smart contract.")

		// Now create a query
		q := mpc.Query{
			QueryType:       mpc.QueryType(int(event.QueryType.Int64())),
			Identifier:      int(event.Identifier.Int64()),
			Attribute:       int(event.Attribute.Int64()),
			ClientReference: event.ClientReference,
		}

		log.Printf("(SmartContract) q: %v\n", q)

		queryChan <- q

	}
}
