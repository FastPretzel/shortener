generate-grpc-gateway:
	protoc -I ./api/proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./api/proto/shortener/shortener.proto

postgresql:
	docker-compose --profile postgresql up --build

memory:
	docker-compose --profile memory up --build

test:
	@echo "Testing usecase:"
	go test -cover -v ./internal/usecase/
	@echo "Testing repo:"
	go test -cover -v ./internal/repo/
