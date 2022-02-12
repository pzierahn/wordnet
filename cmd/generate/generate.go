package main

import (
	"github.com/pzierahn/wordnet/generate"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	generate.Lexicon()
}
