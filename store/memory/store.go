package memory

import (
	"errors"
	"github.com/cbwfree/micro-game/utils/dtype"
	"sync"
	"time"
)

var (
	// ErrNotFound is returned when a Read key doesn't exist
	ErrNotFound = errors.New("not found")
)

// 内存数据存储
type Store struct {
	sync.RWMutex
	values map[string]*Record
}

func (ms *Store) List() ([]*Record, error) {
	ms.RLock()
	defer ms.RUnlock()

	var values []*Record

	for _, v := range ms.values {
		if v.CheckState() {
			values = append(values, v)
		}
	}

	return values, nil
}

func (ms *Store) Read(keys ...string) ([]*Record, error) {
	ms.RLock()
	defer ms.RUnlock()

	var records []*Record

	for _, key := range keys {
		v, ok := ms.values[key]
		if !ok {
			return nil, ErrNotFound
		}

		if !v.CheckState() {
			return nil, ErrNotFound
		}

		records = append(records, v)
	}

	return records, nil
}

func (ms *Store) Write(records ...*Record) error {
	ms.Lock()
	defer ms.Unlock()

	for _, r := range records {
		ms.values[r.key] = r
	}

	return nil
}

func (ms *Store) Delete(keys ...string) error {
	ms.Lock()
	defer ms.Unlock()

	for _, key := range keys {
		delete(ms.values, key)
	}

	return nil
}

func (ms *Store) Get(key string) (interface{}, error) {
	ms.RLock()
	defer ms.RUnlock()

	v, ok := ms.values[key]
	if !ok {
		return nil, ErrNotFound
	}

	return v.Value(), nil
}

func (ms *Store) Set(key string, value interface{}, expiry ...time.Duration) error {
	return ms.Write(NewRecord(key, value, expiry...))
}

func (ms *Store) Int(key string) (int, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return dtype.ParseInt(val), nil
}

func (ms *Store) Int32(key string) (int32, error) {
	val, err := ms.Get(key)
	if err != nil {
		return int32(0), err
	}
	return dtype.ParseInt32(val), nil
}

func (ms *Store) Int64(key string) (int64, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return dtype.ParseInt64(val), nil
}

func (ms *Store) Uint32(key string) (uint32, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return dtype.ParseUint32(val), nil
}

func (ms *Store) Uint64(key string) (uint64, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return dtype.ParseUint64(val), nil
}

func (ms *Store) Bool(key string) (bool, error) {
	val, err := ms.Get(key)
	if err != nil {
		return false, err
	}
	return dtype.ParseBool(val), nil
}

// NewStore returns a new store.Store
func NewStore() *Store {
	return &Store{
		values: make(map[string]*Record),
	}
}
