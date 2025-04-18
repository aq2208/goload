generate:
	protoc -I=. \
		--go_out=./internal/generated \
		--go-grpc_out=./internal/generated \
		--grpc-gateway_out=./internal/generated \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out=./api \
		--openapiv2_opt generate_unbound_methods=true \
		api/proto/goload.proto 