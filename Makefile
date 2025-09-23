.PHONY: build

build:
	sam build -t deploy/template.yaml -b build
deploy:
	sam deploy --guided
local-start:
	sam local start-api -t deploy/template.yaml
local-invoke:
	sam local invoke -t deploy/template.yaml
