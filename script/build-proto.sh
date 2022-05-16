#protoc --go_out=./internal/grpc --go_opt=paths=source_relative --go-grpc_out=./internal/grpc --go-grpc_opt=paths=source_relative api/proto.proto

protoc -I/usr/local/include -I. \
   --grpc-gateway_out=./internal/grpc \
   --go_out=./internal/grpc --go_opt=paths=source_relative --go-grpc_out=./internal/grpc --go-grpc_opt=paths=source_relative \
  api/proto.proto

