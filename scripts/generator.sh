#!/bin/sh
rm -rf api docs *.go
docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli:v5.3.1 generate \
    -i /local/openapi.yaml \
    -g go \
    -o /local \
    --git-user-id qovery \
    --git-repo-id qovery-client-go \
    --package-name qovery \
    --skip-validate-spec
gofmt -s -w .