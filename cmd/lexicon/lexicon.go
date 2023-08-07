package main

import (
	"encoding/json"
	"github.com/pzierahn/wordnet/lexicon"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	start := time.Now()

	if len(os.Args) < 2 {
		log.Fatalln("missing word")
	}

	word := os.Args[1]

	entry := lexicon.Get(word)
	byt, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		log.Fatalf("cloudn't marshal json: %v", err)
	}

	log.Printf("time=%v entry=%s", time.Now().Sub(start), byt)
}
