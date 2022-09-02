# About

Benchmarks to compare json, bson and protocol buffers Marshall and Unmarshall

## Installation

Install `protoc` the protocol buffer compiler via [downloads](https://developers.google.com/protocol-buffers/docs/downloads):

or via Nix:

```shell
nix-env -iA nixpkgs.protobuf3_11
```

Install the Go protocol buffers plugin:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Run the compiler to generate the Go struct:

```shell
protoc --go_out=. student.proto 
```

## Run the benchmarks

```shell
cd benchmark (maybe skip this)
go test -run none -bench . -benchmem -benchtime=10s
```

## Resources
- https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
- https://developers.google.com/protocol-buffers/docs/gotutorial
