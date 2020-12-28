module github.com/cbwfree/micro-game/example

go 1.15

require (
	github.com/cbwfree/micro-game v1.3.0
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.23.0
)

replace (
	github.com/cbwfree/micro-game => ../
)