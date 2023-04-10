.PHONY: clean build test run swagger serve-swagger

BINARY_NAME=bin/shortdesc

clean:
	go clean
	rm ${BINARY_NAME}

build:
	go build -o ${BINARY_NAME} *.go

test:
	go test ./... -v

run:
	go run .

swagger:
	swagger generate spec -o docs/swagger.yaml --scan-models

serve-swagger:
	swagger serve --flavor=swagger docs/swagger.yaml
