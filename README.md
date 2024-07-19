# Vast Data Client

This repository contains an auto-generated Golang SDK for interacting with the Vast Data API.

## How to Use This Repository

The [Makefile](Makefile) for this repository includes a target that uses the [OpenAPI
generator](https://github.com/OpenAPITools/openapi-generator) to generate a Golang SDK from the [Vast Data API
spec](vast_swagger_api.yaml) provided by Vast. To regenerate the client after making a change to the spec, run:

```
make vast-client
```
