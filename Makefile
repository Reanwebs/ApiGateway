proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/auth/proto/*.proto
wire:
	cd pkg/di/ && wire

run:
	go run cmd/main.go