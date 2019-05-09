package main

import (
	"github.com/odysseyhack/mpan-compute-initiator/mpc"
	"github.com/odysseyhack/mpan-compute-initiator/smartcontract"
)

func main() {
	// Start our MPC interface
	qc := mpc.StartQueryListener()
	smartcontract.WaitForQuery(qc)
}
