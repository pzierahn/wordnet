package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func UnmarshalGob(byt []byte, obj interface{}) (err error) {
	buf := bytes.NewBuffer(byt)
	dec := gob.NewDecoder(buf)

	return dec.Decode(obj)
}

func UnmarshalJSON(byt []byte, obj interface{}) (err error) {
	buf := bytes.NewBuffer(byt)
	return json.NewDecoder(buf).Decode(&obj)
}
