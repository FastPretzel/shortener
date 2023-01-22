generate-grpc-gateway:
	protoc -I ./api/proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./api/proto/shortener/shortener.proto

postgresql:
	docker-compose --profile postgresql up --build

memory:
	docker-compose --profile memory up --build
