# This repo is for studying `gRPC Up & Running`

## Before build server

You should install some proto and grpc related plugins.
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

# It might need to update your PATH
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

Make it sure you can generate go.pb file with `grpc` flag.
```shell
# Proper command
$ protoc -I ecommerce ecommerce/proto_info.proto --go_out=plugins=grpc:./ecommerce

# If you generate go.pb file with below command you should handle two files.
$ protoc --go_out=./ecommerce --go-grpc_out=./ecommerce ecommerce/proto_info.proto

# With this command you cannot get grpc related functions.
$ protoc --go_out=./ecommerce ecommerce/proto_info.proto 

```

## Note
If you use Goland, you can meet this error.

```
Cannot use 'ms' (type *messageState) as the type protoreflect.Message
Type does not implement 'protoreflect.Message'
need the method: ProtoMethods() *methods 
have the method: ProtoMethods() *protoiface.Methods
```

But, it is only IDE side issue, you can ignore it.
https://github.com/grpc/grpc-go/issues/4980