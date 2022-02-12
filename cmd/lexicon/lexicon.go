package main

import (
	"encoding/json"
	"log"
	"time"
	"wordnet/lexicon"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	start := time.Now()

	entry := lexicon.Get("adjust")
	byt, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		log.Fatalf("cloudn't marshal json: %v", err)
	}

	log.Printf("time=%v entry=%s", time.Now().Sub(start), byt)
}
