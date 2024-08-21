# Go First Steps
## Commands
- `go version`
  - `go version go1.20.5 darwin/arm64`
  - `go version go1.20.5 linux/amd64`
- Format all files `go fmt ./...`
- Scans for common coding mistakes `go vet ./...`
- `go build main.go`
- `go run main.go`
- `go mod init mod_name`
- `go get tidy`

## Makefiles
It lets developers specify a set of operations that are necessary to build a program and the order in which the steps 
must be performed. You may not be familiar with make, but it’s been used to build programs on Unix systems since 1976.