FUNCTIONS = $(shell  find lambda/* -type d -exec basename {} \;)
GITHUB_SHA = $(shell git rev-parse HEAD)

install_deps:
	go get -u github.com/aws/aws-lambda-go/lambda
	go get -u golang.org/x/lint/golint

build_all:
	$(foreach f, $(FUNCTIONS), $(shell GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/$(f)/main lambda/$(f)/main.go))

test_all:
	go test -v ./...

zip_functions:
	find build/* -type d -execdir zip -r '{}-${GITHUB_SHA}.zip' '{}' \;

