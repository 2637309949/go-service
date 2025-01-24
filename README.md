## Install

First install protoc. You can do this with either your package manager, or [directly](https://github.com/protocolbuffers/protobuf/releases) by downloading `protoc-$VERSION-$PLATFORM.zip`

Then install `protoc-gen-go` and `protoc-gen-micro`.

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
cd ./comm/cmd/protoc-gen-micro && go install && cd ../../../
```

Also required: 

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://google.golang.org/protobuf)

protoc --plugin=protoc-gen-go=$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$GOPATH/bin/protoc-gen-micro --proto_path=. --micro_out=. --go_out=. greeter.proto