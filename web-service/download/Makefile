FILES= $(wildcard ../proto/web/*.proto)
export PATH := $(PATH):$(GOPATH)/bin

define compile_proto
@for file in $(FILES); do \
	(echo -e "protoc $$file" && \
	protoc -I../comm/include --plugin=protoc-gen-web=$$GOPATH/bin/protoc-gen-micro --proto_path=../ --web_out=plugins=web:../proto/web $$file) \
	done;
sed -i -e 's/,omitempty//g' ../proto/web/*.go
endef

build:
	$(call compile_proto)
