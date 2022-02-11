package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"wordnet/generate"
	"wordnet/parse"
)

const exportDir = "metadata"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func exportDictionary() {
	words := make(map[string]bool)

	index := parse.Index()
	for _, entry := range index {
		words[entry.Lemma] = true
	}

	exceptions := parse.Exc()
	for _, exc := range exceptions {
		words[exc] = true
	}

	var dictionary []string

	for word := range words {
		dictionary = append(dictionary, word)
	}

	sort.Strings(dictionary)

	byt, err := json.MarshalIndent(dictionary, "", "  ")
	if err != nil {
		log.Fatalf("cloudn't marshal json: %v", err)
	}

	_ = os.MkdirAll(exportDir, 0755)
	err = ioutil.WriteFile(exportDir+"/dictionary.json", byt, 0644)
	if err != nil {
		log.Fatalf("cloudn't write file: %v", err)
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(dictionary); err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(exportDir+"/dictionary.gob", buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("cloudn't write: %v", err)
	}
}

func main() {

	//exportDictionary()

	generate.Lexicon()
}
