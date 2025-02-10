FILES= $(wildcard ../proto/user/*.proto)
export PATH := $(PATH):$(GOPATH)/bin

define compile_proto
@for file in $(FILES); do \
	(echo -e "protoc $$file" && \
	protoc -I../comm/include --plugin=protoc-gen-go=$$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$$GOPATH/bin/protoc-gen-micro --proto_path=../ --micro_out=../proto/user --go_out=../proto/user $$file) \
	done;
sed -i -e 's/,omitempty//g' ../proto/user/*.go
endef

build:
	$(call compile_proto)
