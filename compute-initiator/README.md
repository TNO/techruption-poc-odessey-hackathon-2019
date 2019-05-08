# Compute initiator

Interface between the blockchain and the MPC nodes.

Listens for events in the BC and passes them to an MPC node. The event's parameters are written to a named pipe (aka fifo) to be read by the node. Thus, each node should have its own instance of this program running. Likewise, it is assumed that a `geth` ethereum node is running locally. See the PPA project for documentation on how to accomplish that.

## Usage

```bash
mkfifo /path/to/fifo
./main --filename /path/to/fifo
```

Various parameters relating to the ethereum node are set in constants in `main.go`.
