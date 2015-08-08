.PHONY: build doc fmt lint dev test vet godep

build: vet \
	test \
	go build -v -o ./bin/go-gin-mgo-demo

doc:
	godoc -http=:6060

fmt:
	go fmt ./...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./...

dev:
	go get && go install && PORT=7000 DEBUG=* gin -p 9000 -a 7000 -i run

test:
	go test ./...

# https://godoc.org/golang.org/x/tools/cmd/vet
vet:
	go vet ./...

godep:
	godep save ./...
