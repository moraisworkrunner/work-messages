# work-messages

protobuf for work messages

## setup

1. install protoc to your PATH: https://github.com/protocolbuffers/protobuf/releases

1. install the protoc plugin for your language. For Go, run the go get command to install the protoc plugin for Go:

```shell
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
```

## regenerate definitions

```shell
protoc --go_out=. --go_opt=paths=source_relative work.proto
```
