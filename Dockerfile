FROM golang:1.8-alpine
RUN apk add --update git
RUN apk add --update make
RUN apk add --update protobuf-dev
RUN apk add --update ca-certificates
RUN go get google.golang.org/grpc
RUN go get google.golang.org/grpc
RUN go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/orcaman/concurrent-map
RUN go get github.com/petar/GoLLRB/llrb
RUN go get github.com/satori/go.uuid
COPY . /go/src/github.com/emef/tally
WORKDIR /go/src/github.com/emef/tally
RUN make build install