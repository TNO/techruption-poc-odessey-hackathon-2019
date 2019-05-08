# techruption-multi-party-all-night

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

A separate blockchain should be run. Refer to PPA documentation and the smart contracts repo for the hackathon.
