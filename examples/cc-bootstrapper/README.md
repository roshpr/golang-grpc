go mod init roshpr.net/example/cc-bootstrapper

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

``
$  protoc --go_out=bootstrap --go_opt=paths=source_relative  --go-grpc_out=bootstrap --go-grpc_opt=paths=source_relative  bootstrap.proto
``
go mod tidy 

Add the below options to generate openapi swagger docs
``
--openapiv2_out bootstrap  --openapiv2_opt logtostderr=true --openapiv2_opt generate_unbound_methods=true
``
