#!/bin/bash

set -ex

docker stack deploy -c ./whatsmyip.stack.yml whatsmyip
