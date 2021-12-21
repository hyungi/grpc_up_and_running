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
# In text book, it uses this command
$ protoc -I ecommerce ecommerce/proto_info.proto --go_out=plugins=grpc:./ecommerce

# But, after grpc version update, plugins keyword is not working anymore.
# So, if you use latest version(after v1.2 protoc-gen-go) of grpc please use this command and I will use it, too. 
$ protoc -I ecommerce ecommerce/proto_info.proto --go_out=./ecommerce --go-grpc_out=./ecommerce

# With this command you cannot get grpc related functions.
# DO NOT USE IT
$ protoc --go_out=./ecommerce ecommerce/proto_info.proto 
```
refer: https://stackoverflow.com/questions/61044883/switch-from-go-out-plugins-to-go-grpc-out-path-problem

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