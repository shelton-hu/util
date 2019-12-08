package jsonutil

import (
	"encoding/json"

	"github.com/bitly/go-simplejson"
)

func Encode(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

func Decode(y []byte, o interface{}) error {
	return json.Unmarshal(y, o)
}

func ToString(o interface{}) string {
	b, err := Encode(o)
	if err != nil {
		return ""
	}
	return string(b)
}

// FIXME: need improve performance
func ToJson(o interface{}) (Json, error) {
	var j Json
	j = &fakeJson{simplejson.New()}
	b, err := Encode(o)
	if err != nil {
		return j, nil
	}
	j, err = NewJson(b)
	if err != nil {
		return nil, err
	}
	return j, nil
}

type fakeJson struct {
	*simplejson.Json
}

func NewJson(y []byte) (Json, error) {
	j, err := simplejson.NewJson(y)
	return &fakeJson{j}, err
}

func (j *fakeJson) Get(key string) Json {
	return &fakeJson{j.Json.Get(key)}
}

func (j *fakeJson) GetPath(branch ...string) Json {
	return &fakeJson{j.Json.GetPath(branch...)}
}

func (j *fakeJson) CheckGet(key string) (Json, bool) {
	result, ok := j.Json.CheckGet(key)
	return &fakeJson{result}, ok
}
