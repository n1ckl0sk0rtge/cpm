build:
	go build -tags dev
run:
	go run cpm.go
release:
	go build
testenv:
	docker build -f tests/testEnvironment.Dockerfile -t cpm-testenvironment:latest .
test:
	docker run --privileged --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v ${PWD}:/app cpm-testenvironment:latest