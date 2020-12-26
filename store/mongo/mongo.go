// MongoDB 连接
// * 查找单个文档时, 如果未找到文件, 则会返回 ErrNoDocuments 错误
// * 查找多个文档时, 如果未找到任何文档, 则会返回 ErrNilDocument 错误
// * bson.M 是无序的 doc 描述
// * bson.D 是有序的 doc 描述
// * bsonx.Doc 是类型安全的 doc 描述
package mongo

import (
	"context"
	"errors"
	"github.com/cbwfree/micro-game/utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
	"time"
)

// MongoDB 数据存储
type Store struct {
	opts   *Options
	ctx    context.Context
	cancel context.CancelFunc
	client *mongo.Client
}

func (s *Store) Init(opts ...Option) {
	for _, o := range opts {
		o(s.opts)
	}
}

func (s *Store) DbName() string {
	return s.opts.DbName
}

func (s *Store) Opts() *Options {
	return s.opts
}

func (s *Store) Connect() error {
	if s.client != nil {
		return nil
	}

	if s.opts.RawUrl == "" {
		s.opts.RawUrl = "mongodb://127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/?replicaSet=rs1"
	} else if !strings.HasPrefix(s.opts.RawUrl, "mongodb://") {
		s.opts.RawUrl = "mongodb://" + s.opts.RawUrl
	}

	opts := options.Client().
		SetMinPoolSize(s.opts.MinPoolSize).
		SetMaxPoolSize(s.opts.MaxPoolSize).
		SetConnectTimeout(s.opts.ConnectTimeout).
		SetSocketTimeout(s.opts.SocketTimeout).
		SetMaxConnIdleTime(s.opts.MaxConnIdleTime).
		SetRetryWrites(true).
		SetRetryReads(true).
		ApplyURI(s.opts.RawUrl)
	if opts.ReplicaSet == nil || *opts.ReplicaSet == "" {
		return errors.New("this system only supports replica sets. example: mongodb://0.0.0.0:27017,0.0.0.0:27018,0.0.0.0:27019/?replicaSet=rs1")
	}

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer s.cancel()

	if mc, err := mongo.Connect(s.ctx, opts); err != nil {
		return err
	} else {
		s.client = mc
	}

	// 检查MongoDB连接
	if err := s.client.Ping(s.ctx, readpref.Primary()); err != nil {
		return err
	}

	log.Debug("Store [mongodb] Connect to %s", s.opts.RawUrl)

	return nil
}

func (s *Store) Disconnect() error {
	if s.client == nil {
		return nil
	}

	if err := s.client.Disconnect(s.ctx); err != nil {
		return err
	}

	s.client = nil

	return nil
}

// Client 获取客户端
func (s *Store) Client() *mongo.Client {
	return s.client
}

// Database 获取数据库对象
func (s *Store) D(dbname ...string) *mongo.Database {
	if len(dbname) > 0 && dbname[0] != "" {
		return s.client.Database(dbname[0])
	}
	return s.client.Database(s.opts.DbName)
}

// Collection 获取集合对象
func (s *Store) C(name string, dbname ...string) *mongo.Collection {
	if len(dbname) > 0 && dbname[0] != "" {
		return s.client.Database(dbname[0]).Collection(name)
	}
	return s.client.Database(s.opts.DbName).Collection(name)
}

// CloneCollection 克隆集合对象
func (s *Store) CloneC(name string, dbname ...string) (*mongo.Collection, error) {
	return s.C(name, dbname...).Clone()
}

func (s *Store) GetIncId(id string) (int64, error) {
	return GetIncId(context.Background(), s.D(), id)
}

// 获取集合列表
func (s *Store) ListCollectionNames(dbname ...string) ([]string, error) {
	return s.D(dbname...).ListCollectionNames(context.Background(), bson.M{})
}

func (s *Store) Scan(dbName, tabName string, cur, size int64, filter interface{}, result interface{}, fn ...func(opts *options.FindOptions) *options.FindOptions) *Scan {
	var scan *Scan
	_ = s.Client().UseSession(context.Background(), func(sctx mongo.SessionContext) error {
		col := sctx.Client().Database(dbName).Collection(tabName)

		count, _ := col.CountDocuments(sctx, filter)
		scan = NewScan(count, cur, size)

		if count > 0 {
			opts := scan.FindOptions()
			if len(fn) > 0 {
				opts = fn[0](opts)
			}
			cur, err := col.Find(sctx, filter, opts)
			if err != nil {
				return err
			}
			if err := cur.All(nil, result); err != nil {
				return err
			}
		}

		return nil
	})
	return scan
}

func NewStore(opts ...Option) *Store {
	ms := &Store{
		opts: newOptions(opts...),
	}
	return ms
}
