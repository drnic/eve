#!/bin/bash

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
cd $ROOT

set -eu

mkdir -p tmp
rm -rf tmp/bosh-scaling-operator.yml

echo "convert to local file"
go run main.go convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge' \
  --target tmp/bosh-scaling-operator.yml


echo "convert to stdout"
go run main.go convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge'
