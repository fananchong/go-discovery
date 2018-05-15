#!/bin/bash

set -ex

docker stack deploy -c ./docker-stack-etcd-static.yml etcd-static
