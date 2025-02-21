FILES= $(wildcard ../proto/test/*.proto)
export PATH := $(PATH):$(GOPATH)/bin

define compile_proto
@for file in $(FILES); do \
	(echo -e "protoc $$file" && \
	protoc -I../comm/include --plugin=protoc-gen-go=$$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$$GOPATH/bin/protoc-gen-micro --proto_path=../ --micro_out=../proto/test --go_out=../proto/test $$file) \
	done;
sed -i -e 's/,omitempty//g' ../proto/test/*.go
endef

build:
	$(call compile_proto)
