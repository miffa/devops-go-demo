all: test devops-go-demo
 
# Run tests
test: 
	go test ./... 
 
# Build binary
devops-go-demo: 
	go build -o devops-go-demo main.go
