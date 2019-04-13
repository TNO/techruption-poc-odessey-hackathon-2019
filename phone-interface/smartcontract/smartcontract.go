package smartcontract

//go:generate abigen -abi gate_keeper.abi -pkg smartcontract -type ComputeContract -out gatekeeper.go

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/odysseyhack/mpan-compute-initiator/mpc"
)

const (
	SMARTCONTRACT_ADDRESS = "0x368f79382cc5a7b769134369a2de7f5b97b28041"
	ETHEREUM_URL          = "ws://134.221.210.18:8546"
)

func WaitForQueries() chan mpc.Query {
	ch := make(chan mpc.Query)
	conn, err := ethclient.Dial(ETHEREUM_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	gatekeeper, err := NewGovernance(common.HexToAddress(SMARTCONTRACT_ADDRESS), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate the Governance contract: %v", err)
	}
	computeAddress, err := gatekeeper.GetContractAddress()
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

// temporary
type ComputeContract struct{}

func NewComputeContract(address common.Address, backend bind.ContractBackend) (*ComputeContract, error) {
	return nil, nil
}

func (_ComputeContract *ComputeContract) SubmitQuery(q *mpc.Query) error {
	return nil
}

func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	return nil, nil
}

type Governance struct{}

func (_Governance *Governance) GetContractAddress() (common.Address, error) {
	return common.HexToAddress("0x0"), nil
}
