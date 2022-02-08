package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"wordnet/assets"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	records := assets.IndexPOS()

	terms := make(map[string]bool)

	for _, record := range records {
		terms[record[0]] = true
	}

	log.Printf("%d terms", len(terms))

	var dict []string
	for term := range terms {
		dict = append(dict, term)
	}

	sort.Strings(dict)

	dictFile := filepath.Join("dumb", "dictionary.en.json")
	dump, _ := json.MarshalIndent(dict, "", "  ")
	err := ioutil.WriteFile(dictFile, dump, 0755)
	if err != nil {
		log.Fatalf("cloudn't write %s", dictFile)
	}
}
