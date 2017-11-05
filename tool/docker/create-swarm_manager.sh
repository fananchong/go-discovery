#!/bin/bash

if [ ! $1 ]; then
    echo "./_create-swarm_manager.sh ip"
fi


set -e

node_ip=$1

docker swarm init --advertise-addr=$node_ip


