package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	single *Store
	once   sync.Once
)

func S() *Store {
	once.Do(func() {
		single = NewStore()
	})
	return single
}

func Connect() error {
	return S().Connect()
}

func Disconnect() error {
	return S().Disconnect()
}

// Client 获取客户端
func Client() *mongo.Client {
	return S().client
}

func DB(dbname ...string) *mongo.Database {
	return S().D(dbname...)
}

func Col(name string, dbname ...string) *mongo.Collection {
	return S().C(name, dbname...)
}
