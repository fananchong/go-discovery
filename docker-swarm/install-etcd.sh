#!/bin/bash


if [ $# != 2 ] ; then
    echo "USAGE: $0 IP TOKEN"
    exit
fi

set -e

cp -f docker-stack-etcd.yml docker-stack-etcd.yml.temp
sed 's/IP/'$1'/g' docker-stack-etcd.yml.temp
sed 's/TOKEN/'$2'/g' docker-stack-etcd.yml.temp
docker stack deploy -c ./docker-stack-etcd.yml.temp etcd
rm -rf ./docker-stack-etcd.yml.temp


