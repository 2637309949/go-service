FILES= $(wildcard ../proto/cache/*.proto)
export PATH := $(PATH):$(GOPATH)/bin

define compile_proto
@for file in $(FILES); do \
	(echo -e "protoc $$file" && \
	protoc -I../comm/include --plugin=protoc-gen-go=$$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$$GOPATH/bin/protoc-gen-micro --proto_path=../ --micro_out=../proto/cache --go_out=../proto/cache $$file) \
	done;
sed -i -e 's/,omitempty//g' ../proto/cache/*.go
endef

build:
	$(call compile_proto)
