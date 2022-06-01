build:
	go build -tags dev
run:
	go run cpm.go
release:
	go build
test:
	go test ./...