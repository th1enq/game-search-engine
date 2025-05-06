proto:
	protoc -I=proto -I=googleapis \
		--go_out=proto \
		--go-grpc_out=proto \
		--grpc-gateway_out=proto \
		--openapiv2_out=doc \
		proto/game.proto
