#!/bin/bash

set -ex

docker stack deploy -c ./docker-stack-discovery_etcd_io.yml discovery_etcd_io
