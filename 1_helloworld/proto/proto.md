
- url : https://github.com/micro/protoc-gen-micro

```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```