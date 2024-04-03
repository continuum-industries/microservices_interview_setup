FROM golang:1.16.0

RUN apt-get update && apt install -y protobuf-compiler

# # Get the source from GitHub
RUN go get google.golang.org/grpc@v1.36.0
# # Install protoc-gen-go
RUN go get github.com/golang/protobuf/proto@v1.4.2
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

# Set necessary environmet variables needed for our image
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY ./gateway-api/src/go.mod .
COPY ./gateway-api/src/go.sum .
RUN go mod download

# Copy the code into the container

COPY ./gateway-api ./gateway-api
COPY ./shared/protobuffers/ ./shared/protobuffers/

WORKDIR /build/gateway-api

RUN make build_proto

# Build the application

WORKDIR /build/gateway-api/src

RUN go build -o /build/main .

# Copy binary from build to main folder
RUN cp /build/main .

ENTRYPOINT [ "go", "run", "." ]
