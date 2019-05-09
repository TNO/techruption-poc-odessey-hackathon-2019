#!/bin/sh

# start the compute initiators (blockchain (-> CI -> mpc node) x5)

/home/mpc/compute-initiator/main -filename /tmp/nodepipe1 > /tmp/compute-initiator-log1 &
/home/mpc/compute-initiator/main -filename /tmp/nodepipe2 > /tmp/compute-initiator-log2 &
/home/mpc/compute-initiator/main -filename /tmp/nodepipe3 > /tmp/compute-initiator-log3 &
/home/mpc/compute-initiator/main -filename /tmp/nodepipe4 > /tmp/compute-initiator-log4 &
/home/mpc/compute-initiator/main -filename /tmp/nodepipe5 > /tmp/compute-initiator-log5 &

/home/mpc/gabriele/node.py /tmp/nodepipe1&
/home/mpc/gabriele/node.py /tmp/nodepipe2&
/home/mpc/gabriele/node.py /tmp/nodepipe3&
/home/mpc/gabriele/node.py /tmp/nodepipe4&
/home/mpc/gabriele/node.py /tmp/nodepipe5&
