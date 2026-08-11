start/cli:
	go run example/cli/main.go

start/holidays:
	go run example/holidays/main.go

start/http:
	go run example/http/main.go

tools:
	which golint # go install golang.org/x/lint/golint@latest
	which staticcheck # go install honnef.co/go/tools/cmd/staticcheck@latest

lint:
	gofmt  -w=true -s=true -l=true ./
	golint ./...
	go vet ./...
	staticcheck ./...
