package byteutil

import "encoding/json"

type Bytes struct {
	bytes []byte
}

func Write(bytes []byte) *Bytes {
	return &Bytes{
		bytes: bytes,
	}
}

func WriteObject(obj interface{}) *Bytes {
	bytes, err := json.Marshal(obj)
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
	if err := json.Unmarshal(b.bytes, &obj); err != nil {
		panic(err)
	}
}
