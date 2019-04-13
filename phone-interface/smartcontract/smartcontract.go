package smartcontract

//go:generate abigen -abi governance.abi -pkg smartcontract -type Governance -out gatekeeper.go

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

func WaitForQueries() chan mpc.Query {
	ch := make(chan mpc.Query)
	conn, err := ethclient.Dial(ETHEREUM_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	governance, err := NewGovernance(common.HexToAddress(SMARTCONTRACT_ADDRESS), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate the Governance contract: %v", err)
	}
	computeAddress, err := governance.GetGatekeeperAddress(nil)
	if err != nil {
		log.Fatalf("Failed to talk to the Governance contract: %v", err)
	}

	// Instantiate the contract and display its name
	computeContract, err := NewComputeContract(computeAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate the Compute contract: %v", err)
	}

	go func() {
		for {
			q := <-ch
			computeContract.SubmitQuery(&q)
		}
	}()

	return ch
}

type ComputeContract struct{}

func NewComputeContract(address common.Address, backend bind.ContractBackend) (*ComputeContract, error) {
	return nil, nil
}

func (_ComputeContract *ComputeContract) SubmitQuery(q *mpc.Query) error {
	return nil
}
