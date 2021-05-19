package byteutil

import "github.com/vmihailenco/msgpack"

type Bytes struct {
	bytes []byte
}

func Write(bytes []byte) *Bytes {
	return &Bytes{
		bytes: bytes,
	}
}

func WriteObject(obj interface{}) *Bytes {
	bytes, err := msgpack.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return &Bytes{
		bytes: bytes,
	}
}

func (b *Bytes) Read() []byte {
	return b.bytes
}

func (b *Bytes) ReadObject(obj interface{}) {
	if err := msgpack.Unmarshal(b.bytes, &obj); err != nil {
		panic(err)
	}
}
