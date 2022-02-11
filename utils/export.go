package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"log"
)

func ExportJson(filename string, data interface{}) {
	byt, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("cloudn't marshal data: %v", err)
	}

	err = ioutil.WriteFile(filename, byt, 0644)
	if err != nil {
		log.Fatalf("cloudn't write file: %v", err)
	}
}

func ExportGob(filename string, data interface{}) {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(data); err != nil {
		log.Fatalf("cloudn't gob data: %v", err)
	}

	err := ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("cloudn't write file: %v", err)
	}
}
