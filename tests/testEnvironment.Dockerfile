FROM docker:latest

COPY --from=golang:alpine /usr/local/go/ /usr/local/go/

ENV GOPATH /go
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH $GOPATH/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin &&\
    chmod -R 777 "$GOPATH" &&\
    apk add --no-cache make build-base

WORKDIR /app

ENTRYPOINT go mod download && go test -v ./...