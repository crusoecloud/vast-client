GO_SWAGGER_VERSION = v0.28.0

# This is how we do it in rest gateway and how we 
# .PHONY: update-swagger
# update-swagger: install-swagger 
# 	@swagger generate client --spec=vast_swagger_api.yaml --target=swagger
#
# .PHONY: install-swagger
# install-swagger:
# 	@go install github.com/go-swagger/go-swagger/cmd/swagger@${GO_SWAGGER_VERSION}
#
.PHONY: swagger 
swagger: 
	@swagger-codegen generate -i ./vast_swagger_api.yaml --lang go  -o swagger
	@rm -rf swagger/{docs,api,.gitignore,.swagger-codegen-ignore,.swagger-codegen,.travis.yml,README.md,git_push.sh}
