#!/bin/bash
# start test-network with couchdb as database

# import constants
. constants.sh

./network.sh down && ./network.sh up createChannel -c ${channelname} -ca -s couchdb

sleep 2

./packagedeploy.sh