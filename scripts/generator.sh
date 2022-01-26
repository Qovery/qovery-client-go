
#!/bin/sh
docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli:v5.3.1 generate \
    -i /local/swagger.json \
    -g go \
    -o /local \
    --package-name qovery \
    --skip-validate-spec