package redis

import (
	"fmt"
	"testing"
)

func TestStore_Connect(t *testing.T) {
	rs := NewStore(
		WithDb(0),
		WithUrl("127.0.0.1:6379"),
		WithMinIdleConns(40),
		WithPoolSize(80),
	)
	if err := rs.Connect(); err != nil {
		panic(err)
	}
	defer rs.Disconnect()

	fmt.Printf("redis连接成功\n")
}
