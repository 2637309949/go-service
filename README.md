## Install

First install protoc. 

```
wget https://github.com/protocolbuffers/protobuf/releases/download/v25.6/protoc-25.6-linux-x86_64.zip
unzip protoc-25.6-linux-x86_64.zip -d protoc-25.6-linux-x86_64
sudo mv protoc-25.6-linux-x86_64/bin/protoc /usr/local/bin/
```

Then install `protoc-gen-go` and `protoc-gen-micro`.

```
cd ./comm/cmd/protoc-gen-micro && go install && cd ../../../
cd ./comm/cmd/protobuf-go/cmd/protoc-gen-go && go install && cd ../../../../../
```
