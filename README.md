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
cd ./comm/cmd/protobuf-go/cmd/protoc-gen-validate && go install && cd ../../../../../
```

Then start consul

```
wget https://releases.hashicorp.com/consul/1.20.2/consul_1.20.2_linux_amd64.zip
unzip consul_1.20.2_linux_amd64.zip
./consul agent -server -bootstrap-expect=1 -http-port=8400 -data-dir=comm/config/data -bind="172.30.10.72" -advertise="172.30.10.72" -ui=true

```

## New Gateway Features

### `http2rpc`

Enables the conversion of HTTP requests into RPC calls.

### `http2web`

Facilitates the transformation of HTTP requests into Web requests.

### `http2file`

Allows HTTP requests to be mapped to static file responses.


## New proto Features
### `Expose rpc`

```proto
syntax = "proto3";

option go_package = "/;user";

import "google/api/annotations.proto";
import "proto/user/user.proto";

service UserService {
	rpc QueryUserDetail(UserFilter) returns (User) {
		option (google.api.http) = {
			get: "/QueryUserDetail" // expose httpApi
			body: "*"
		};
	}
	rpc InsertUser(User) returns (User) {} // only rpc
}
```

### `Expose web`

```proto
syntax = "proto3";

option go_package = "/;web";

import "google/api/annotations.proto";
import "http/http.proto";

service WebService {
	rpc FileUpload (http.ResponseWriter) returns (http.Request) {
		option (google.api.http) = {
			post: "/fileupload/upload" // expose httpApi
			authorization: false
			body: "*"
		};
	}
}
```

