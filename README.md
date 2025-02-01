## Install

First install protoc. You can do this with either your package manager, or [directly](https://github.com/protocolbuffers/protobuf/releases) by downloading `protoc-$VERSION-$PLATFORM.zip`

Then install `protoc-gen-go` and `protoc-gen-micro`.

```
cd ./comm/cmd/protoc-gen-micro && go install && cd ../../../
cd ./comm/cmd/protobuf-go/cmd/protoc-gen-go && go install && cd ../../../../../
```

Also required: 

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://google.golang.org/protobuf)

