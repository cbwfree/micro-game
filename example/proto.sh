#!/usr/bin/env bash

# 游戏协议 (不能引用其它包)
protoc --proto_path=./ --go_out=paths=source_relative:. ./proto/msg/*.proto

# 公共模块 (不能含有 service)
protoc --proto_path=./ --go_out=paths=source_relative:. ./proto/com/*.proto

# 独立服务
protoc --proto_path=./ --go_out=paths=source_relative:. --rpc_out=paths=source_relative:. ./proto/gate/*.proto
protoc --proto_path=./ --go_out=paths=source_relative:. --rpc_out=paths=source_relative:. ./proto/game/*.proto
protoc --proto_path=./ --go_out=paths=source_relative:. --rpc_out=paths=source_relative:. ./proto/db/*.proto
protoc --proto_path=./ --go_out=paths=source_relative:. --rpc_out=paths=source_relative:. ./proto/admin/*.proto







