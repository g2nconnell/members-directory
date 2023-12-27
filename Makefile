generate_grpc_code:
	protoc \
	--go_out=directoryService \
	--go_opt=paths=source_relative \
	--go-grpc_out=directoryService \
	--go-grpc_opt=paths=source_relative \
	directoryService.proto
	