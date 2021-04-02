FUNCTIONS = $(shell  find cmd/* -type d -exec basename {} \;)

install_deps:
	go get -u github.com/aws/aws-lambda-go/lambda
	go get -u golang.org/x/lint/golint

build_all:
	$(foreach f, $(FUNCTIONS), $(shell GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/$(f)/main cmd/$(f)/main.go))

zip_functions:
	find build -type d -execdir zip -r '{}.zip' '{}' \;

