package main

import (
	"log"
	"wordnet/generate"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	generate.Lexicon()
}
