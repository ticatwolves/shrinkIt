.PHONY: build

build:
	sam build -t config/template.yaml -b build

deploy:
	sam deploy --guided -t build/template.yaml --config-file `pwd`/config/samconfig.toml

destroy:
	sam delete --config-file `pwd`/config/samconfig.toml

go-build:
	GOOS=linux GOARCH=amd64 go build -o build/bootstrap cmd/main.go

local-start:
	sam local start-api -t config/template.yaml

local-invoke:
	sam local invoke -t config/template.yaml
