proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/common/proto/*.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/common/proto/conference.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/common/proto/videoStreaming.proto

wire:
	cd pkg/common/di/ && wire

run:
	go run cmd/main.go

swag:
	swag init -g ./pkg/api/handlers/user.go -o ./docs
	swag init -g ./pkg/api/handlers/conference.go -o ./docs

docker:
	docker run -p 8080:8080   -e Auth_SRV=:4000   -e Conference_SRV=:5050   -e Video_Service=:8081   -e PORT=:8080   -e JWT_SECRET_KEY=P6zqwuYDJomBGnleDYtF3pMyoN3sVaiy2BTbUNd566g   -e SwagTitle="REAN CONNECT"   -e SwagDescription="Yo Yo Yo 148 3 to the 3 to the 6 to the 9"   -e SwagVersion=1.0   -v "$(pwd)/.env:/app/.env"   -v "$(pwd)/server.crt:/app/server.crt"   my-apigateway-image:new