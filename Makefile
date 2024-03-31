install_dependencies:
	@brew install protobuf

proto: install_dependencies
	@protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto

run: proto
	@go run main.go
