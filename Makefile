.PHONY: build
go-build:
	GOOS=linux GOARCH=amd64 go build -o build/bootstrap cmd/main.go
build:
	sam build -t deploy/template.yaml -b build
deploy:
	sam deploy --guided
local-start:
	sam local start-api -t deploy/template.yaml --docker-network granite
local-invoke:
	sam local invoke -t deploy/template.yaml --docker-network granite
