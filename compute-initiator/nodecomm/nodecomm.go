package nodecomm

import (
	"flag"
	"log"
	"os"
)

var output *os.File

// In init we make a named pipe for comms to the mpc node and prepare it for writing
func init() {
	filename := flag.String("filename", "/tmp/computeInitiatorOutput", "Named pipe for output to mpc node")
	flag.Parse()

	var err error
	output, err = os.OpenFile(*filename, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Fatalf("Can not open named pipe %v, %v", *filename, err)
	}
}

func Send(data []byte) {
	output.Write(data)
}
