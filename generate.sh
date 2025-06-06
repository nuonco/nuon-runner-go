#!/usr/bin/env bash

set -e
set -o pipefail
set -u

env=${NUON_ENV:-dev}
echo >&2 "using $env, please set NUON_ENV to use a different environment."

case $env in
  stage)
    url="https://runner.stage.nuon.co/docs/doc.json"
    ;;
  prod)
    url="https://runner.nuon.co/docs/doc.json"
    ;;
  dev)
    url="http://localhost:8083/docs/doc.json"
    ;;
  *)
    echo "invalid env $env"
    exit 1
    ;;
esac

echo >&2 "generating with OAPI spec from $url"
go run github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5 \
  generate \
  client \
  --skip-tag-packages \
  -f $url

echo >&2 "generating mocks"
go run github.com/golang/mock/mockgen \
  -destination=mock.go \
  -source=client.go \
  -package=nuonrunner
