## Microv5+

基于 `go-micro.dev/v5` 微服务框架深度改造而来的项目，旨在提供更强大、更灵活、开箱即用的企业级解决方案。通过对原框架的全面优化与扩展，引入了全新的网关设计、统一的鉴权机制、高级服务组件，以及一系列贴近实际业务需求的企业级服务构建案例。本框架保留了 go-micro.dev/v5 的轻量与高效特性，同时在功能性、稳定性和易用性上进行了大幅提升，适用于构建现代化、高并发、分布式的企业应用。

## 新增特性

- v5+网关服务
- 支持配置认证，多策略鉴权
- HTTP `protoc` 插件
- PROXY `protoc` 插件
- WEB `protoc` 插件
- 日志Trace
- 默认Mysql，支持分表，db负载
- 跨平台构建脚本
- 企业架构拆箱即用
- 更多特性在文档中描述

## 环境配置

先安装 `protoc`

```bash linenums="1"
hub="https://github.com/protocolbuffers/protobuf"
releases="/releases/download/v25.6/protoc-25.6-linux-x86_64.zip"
wget "$hub$releases"
unzip protoc-25.6-linux-x86_64.zip -d protoc-25.6-linux-x86_64
sudo mv protoc-25.6-linux-x86_64/bin/protoc /usr/local/bin/
```

再安装插件 `protoc-gen-go` and `protoc-gen-micro`.

```bash linenums="1"
cd ./comm/cmd/protoc-gen-micro && go install
cd ./comm/cmd/protobuf-go/cmd/protoc-gen-go && go install
cd ./comm/cmd/protobuf-go/cmd/protoc-gen-validate && go install
```

在Window环境安装自定义make, unix环境不需要
```
cd ./comm/cmd/protobuf-go/cmd/make && go install
```

最后安装配置服务 `consul`
```bash linenums="1"
hub="https://releases.hashicorp.com/consul"
releases="/1.20.2/consul_1.20.2_linux_amd64.zip"
wget "$hub$releases"
unzip consul_1.20.2_linux_amd64.zip
```

启动 `consul`（仅用于测试!）
```bash linenums="1"
./consul agent \
-server \
-bootstrap-expect=1 \
-http-port=8400 \
-data-dir=comm/assert/data \
-bind="172.30.10.72" \
-advertise="172.30.10.72" \
-ui=true
```
