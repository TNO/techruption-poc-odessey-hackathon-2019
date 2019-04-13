package smartcontract

//go:generate abigen -abi governance.abi -pkg smartcontract -type Governance -out gatekeeper.go

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/odysseyhack/mpan-compute-initiator/mpc"
)

const (
	SMARTCONTRACT_ADDRESS = "0xa652605f3794d5cd868aa5f295e60fae924fe836"
	ETHEREUM_URL          = "ws://127.0.0.1:8546"
	KEYFILE               = "/home/mpc/ethernet-node/data/client_node/keystore/UTC--2019-04-13T12-18-34.942247603Z--d33de107e71120c6022465354cedf69810b62678"
	PASSPHRASE            = "mpc_123"
)

func GetGatekeeper() (*Gatekeeper, error) {
	conn, err := ethclient.Dial(ETHEREUM_URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	governance, err := NewGovernance(common.HexToAddress(SMARTCONTRACT_ADDRESS), conn)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate the Governance contract: %v", err)
	}
	computeAddress, err := governance.GetGatekeeperAddress(nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to talk to the Governance contract: %v", err)
	}

	return NewGatekeeper(computeAddress, conn)
}

func GetTxOpts() (*bind.TransactOpts, error) {
	// Prepare keys etc for transactions
	f, err := os.Open(KEYFILE)
	if err != nil {
		log.Fatalf("Failed to open key file: %v", err)
	}
	return bind.NewTransactor(f, PASSPHRASE)
}

func QueryChannel(gatekeeper *Gatekeeper, txOpts *bind.TransactOpts) chan mpc.Query {
	ch := make(chan mpc.Query)
	go func() {
		for {
			q := <-ch

			// Prepare data
			if len(q.ClientReference) > 50 {
				// Limitation in the contract, reference is normally just 36 chars
				q.ClientReference = q.ClientReference[0:50]
			}

			// Send tx
			var tx *types.Transaction
			var err error
			switch q.QueryType {
			case mpc.QUERY_TYPE_INFO:
				tx, err = gatekeeper.SubmitInfoQuery(txOpts, q.ClientReference, big.NewInt(int64(q.Identifier)))

			case mpc.QUERY_TYPE_CALC:
				tx, err = gatekeeper.SubmitCalcQuery(txOpts, q.ClientReference, big.NewInt(int64(q.Identifier)), big.NewInt(int64(q.Attribute)))
			}

			// check result
			if err != nil {
				log.Printf("Error submitting query, %v", err)
			} else {
				log.Printf("Successfully submitted query, tx %+v", tx)
			}
		}
	}()

	return ch
}

func WaitForQueries() chan mpc.Query {
	// Instantiate the contract and display its name
	gatekeeper, err := GetGatekeeper()
	if err != nil {
		log.Fatalf("Failed to instantiate the Compute contract: %v", err)
	}

	txOpts, err := GetTxOpts()
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	return QueryChannel(gatekeeper, txOpts)
}
