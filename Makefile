proto:
	protoc internal/grpc/domain/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/grpc/server/main.go

client: 
	go run cmd/grpc/client/main.go