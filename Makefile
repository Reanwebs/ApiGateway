proto:
<<<<<<< HEAD
	protoc --go_out=. --go-grpc_out=. ./pkg/auth/proto/*.proto
=======
	protoc --go_out=. --go-grpc_out=. ./pkg/auth/common/pb/*.proto
>>>>>>> rohith
wire:
	cd pkg/auth/common/di/ && wire

run:
	go run cmd/main.go