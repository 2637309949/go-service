module github.com/go-micro/generator/cmd/protoc-gen-micro

go 1.22

require (
	google.golang.org/genproto v0.0.0-20221010155953-15ba04fc1c0e
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/google/go-cmp v0.5.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace (
	google.golang.org/genproto => ../../../comm/proto/genproto
	github.com/micro/plugins/v5/registry/consul => ../comm/plugins/registry/consul
	github.com/micro/plugins/v5/wrapper/trace/opentracing => ../comm/plugins/wrapper/trace/opentracing
	go-micro.dev/v5 => ../comm/go-micro
)