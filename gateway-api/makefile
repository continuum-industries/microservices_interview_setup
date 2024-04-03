SHELL=/bin/bash
protobuf_path = ../shared_modules/pareto_common/protobuffers

.PHONY:
	check
	go_pkg_install
	build_proto

check:
	@[[ -f `which go` ]] || { echo -e "\nPlease install go first\n"; exit 1; }
	@[[ -f `which protoc` ]] || { echo -e "\nPlease install protoc first (protobuf-compiler on apt)\n"; exit 1; }

go_pkg_install:
	go get github.com/golang/protobuf/proto@v1.4.2
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

build_proto: check
	@bash deployment/get_protos.sh

test:
	@cd src; go test ./...

run:
	@cd src; go run main.go
