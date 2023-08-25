proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/common/proto/*.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/common/proto/conference.proto

wire:
	cd pkg/common/di/ && wire

run:
	go run cmd/main.go