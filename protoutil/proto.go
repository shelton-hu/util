package protoutil

import (
	"bytes"
	"encoding/json"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

var m = jsonpb.Marshaler{
	EmitDefaults: true,
	OrigName:     true,
	EnumsAsInts:  true,
}

func ParseProto(p proto.Message) (map[string]interface{}, error) {
	var b bytes.Buffer
	if err := m.Marshal(&b, p); err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(b.Bytes(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func ParseProtoToString(p proto.Message) (string, error) {
	return m.MarshalToString(p)
}
