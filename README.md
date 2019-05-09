# Techruption: multi-party all night

This is the result of our participation in the Odyssey Hackathon 2019.

# Summary

In case of an emergency involving a hazardous materials transport, fire brigade workers need accurate information about the contents of the transport and what to do and what not to do. This information is typically spread over multiple commercial parties that may not always be online, yet may also not wish to divulge all information about their transports to a central agency just in case it's needed.

This proof-of-concept addresses this situation by *secret sharing* the database of information about the transports between multiple parties. This way, a fire brigade worker only needs a subset (e.g. half) of the participating companies to be online in order to recover the necessary information, but none of the companies can learn about the others' data individually.

Because of the time-sensitive nature of disaster response, our approach avoids "asking for permission" in favour of auditing afterwards. A blockchain is used to keep an immutable log of requests made to the database, which can be audited independently by the participants.

# Moving parts

A request is handled as follows.

* The fire brigadier uses the phone app to scan a QR code affixed to the transport, which contains its ID. It sends this ID to the *phone interface*.
* The phone interface assigns the request a unique guid, and sends a transaction to the blockchain.
* The blockchain contains two smart contracts to handle this request: a governor, which points to the current version of the gatekeeper, and a gatekeeper, which checks if the person making the request is authorised to do so.
* The gatekeeper emits an event with the request parameters if the request is valid.
* The *compute initiator* monitors the blockchain for this event and sends the request and its ID to the mpc node.
* The *compute on share* nodes retrieve the shares of the requested secret from their database and send it to the combiner.
* The *combiner* takes these shares and retrieves the secret-shared data about the materials transport, and sends this back to the phone interface.
* The phone app keeps a connection open to the phone interface, and receives the information as it arrives.

# Usage

The various parts of the chain should be run as follows, on the same server.

## Phone interface

See its own readme.

## Compute initiator

Four times, `.main --filename fifoX`, `X` being 0 through 3.

## MPC node

Four times, `while :; do python compute_on_share.py X; done`, `X` being 0 through 3.

## MPC combiner

`python combiner.py`

## Blockchain node

`geth --datadir data/client_node --config config/client_config.toml --verbosity 3 console`

# Blockchain

A separate blockchain should be run. Refer to geth documentation and the smart contracts repo for the hackathon.
