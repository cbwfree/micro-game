package mongo

import (
	"fmt"
	"testing"
)

func TestStore_Connect(t *testing.T) {
	rs := NewStore(
		WithDbName("test"),
		WithMinPoolSize(20),
		WithMaxPoolSize(100),
		WithUrl("mongodb://127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/admin?replicaSet=rs1"),
	)
	if err := rs.Connect(); err != nil {
		panic(err)
	}
	defer rs.Disconnect()

	fmt.Printf("MongoDB连接成功\n")
}
