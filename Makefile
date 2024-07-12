.PHONY: vast-client
vast-client:
	@docker run --rm -v $$PWD:/local openapitools/openapi-generator-cli generate -i /local/vast_swagger_api.yaml -g go -o /local/vast-client --skip-validate-spec -p structPrefix=true -p disallowAdditionalPropertiesIfNotPresent=false
