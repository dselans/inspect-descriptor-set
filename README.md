Quick little example of parsing a protobuf descriptor file.

```bash
❯ go run main.go -f assets/test.fds -v
File: common/ps_common_auth.proto
	Package: protos.common
	Syntax: proto3
	Messages:
		Auth
			Fields:
				token (type: TYPE_STRING number: 1)

File: common/ps_common_status.proto
	Package: protos.common
	Syntax: proto3
	Messages:
		Status
			Fields:
				code (type: TYPE_ENUM number: 1)
				message (type: TYPE_STRING number: 2)
				request_id (type: TYPE_STRING number: 3)

-------------- snip ------------------

Found 71 files with 328 messages
```

test.fds was generated via `protoc` against github.com/batchcorp/plumber-schemas/protos,
running the following cmd:

```bash
	docker run --rm -w $(PWD) -v $(PWD):$(PWD) -w${PWD} jaegertracing/protobuf:0.2.0 \
	--proto_path=./protos \
	--proto_path=./protos/args \
	--proto_path=./protos/common \
	--proto_path=./protos/encoding \
	--proto_path=./protos/records \
	--go_out=plugins=grpc:$(GO_PROTOS_DIR) \
	--go_opt=paths=source_relative \
	-o ./$(GO_DESCRIPTOR_SET_DIR)/protos.bin \
	--include_imports \
	--include_source_info \
	protos/*.proto
```
