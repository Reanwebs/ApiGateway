proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/common/proto/*.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/common/proto/conference.proto

wire:
	cd pkg/common/di/ && wire

run:
	go run cmd/main.go

swag:
	swag init -g ./pkg/api/handlers/user.go -o ./docs
	swag init -g ./pkg/api/handlers/conference.go -o ./docs