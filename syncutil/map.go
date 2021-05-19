package syncutil

import (
	"errors"
	"fmt"
	"sync"

	"github.com/vmihailenco/msgpack"
)

type Syncmap struct {
	sync.Map
}

func NewSyncmap() *Syncmap {
	return &Syncmap{
		Map: sync.Map{},
	}
}

func (s *Syncmap) LoadObject(key string, obj interface{}) error {
	value, ok := s.Load(key)
	if !ok {
		return fmt.Errorf("Load key[%s] from syncmap failed", key)
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Value type loaded is not []byte!")
	}

	return msgpack.Unmarshal(bytes, &obj)
}

func (s *Syncmap) LoadOrStoreObject(key string, v interface{}) (loaded bool, err error) {
	bytes, err := msgpack.Marshal(v)
	if err != nil {
		return false, err
	}

	actual, loaded := s.LoadOrStore(key, bytes)

	bytes, ok := actual.([]byte)
	if !ok {
		return false, errors.New("Value type loaded is not []byte!")
	}

	return !loaded, msgpack.Unmarshal(bytes, &v)
}
