FROM golang:1.17.2 as builder

WORKDIR /workdir/go

ADD go.mod ./
ADD cmd ./cmd
ADD internal ./internal
ADD pkg ./pkg

RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 go build -o kubetables cmd/main.go

#_________________________________________________
FROM alpine:3.14

COPY --from=builder /workdir/go/kubetables /usr/local/bin/kubetables

RUN apk add iptables ip6tables

ENTRYPOINT ["/usr/local/bin/kubetables"]