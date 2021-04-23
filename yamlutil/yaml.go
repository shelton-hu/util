package yamlutil

import "github.com/ghodss/yaml"

// Marshals the object into JSON then converts JSON to YAML and returns the YAML.
func Encode(o interface{}) ([]byte, error) {
	return yaml.Marshal(o)
}

// Converts YAML to JSON then uses JSON to unmarshal into an object.
func Decode(y []byte, o interface{}) error {
	return yaml.Unmarshal(y, o)
}
