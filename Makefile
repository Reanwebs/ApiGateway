proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/auth/common/pb/*.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/confernce/common/pb/*.proto

wire:
	cd pkg/auth/common/di/ && wire

run:
	go run cmd/main.go