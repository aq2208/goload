generate:
	protoc -I=. \
		--go_out=./internal/generated/go \
		--go-grpc_out=./internal/generated/grpc \
		--grpc-gateway_out=. \
		--openapiv2_out=./api \
		api/proto/goload.proto