#!/bin/bash

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
cd $ROOT

set -eu

go vet $(go list ./... | grep -v vendor)
go test -v ./...

mkdir -p tmp
rm -rf tmp/bosh-scaling-operator.yml

echo "convert to local file"
go run main.go convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge' \
  --target tmp/bosh-scaling-operator.yml
echo "load values from existing operator file"
values=$(go run main.go values \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --target tmp/bosh-scaling-operator.yml)
expected='{"workers-linux-instances": "5", "workers-linux-instance-type": "m4.xlarge"}'
spruce diff <(echo "$expected") <(echo "$values")

echo "convert to stdout"
go run main.go convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge'

echo "pipe 'convert' to 'values'"
values=$(go run main.go convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge' |
go run main.go values \
  --mapping fixtures/bosh-scaling/mapping.yml)
expected='{"workers-linux-instances": "5", "workers-linux-instance-type": "m4.xlarge"}'
spruce diff <(echo "$expected") <(echo "$values")
