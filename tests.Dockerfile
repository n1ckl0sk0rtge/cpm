FROM docker

COPY --from=golang:alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin &&\
    chmod -R 777 "$GOPATH" &&\
    apk add --no-cache make build-base

COPY . /app

WORKDIR /app

RUN go mod download

RUN go test ./...

