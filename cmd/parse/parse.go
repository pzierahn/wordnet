package main

import (
	"encoding/json"
	"log"
	"wordnet/parse"
)

// https://wordnet.princeton.edu/documentation/wninput5wn
var symbolMapping = map[string]string{
	// various
	"!":  "Antonym",
	"@":  "Hypernym",
	"~":  "Hyponym",
	"=":  "Attribute",
	"+":  "Derivationally related form",
	";c": "Domain of synset - TOPIC",
	";r": "Domain of synset - REGION",
	";u": "Domain of synset - USAGE",
	"^":  "Also see",
	"\\": "Inflection",

	// nouns
	"@i": "Instance Hypernym",
	"~i": "Instance Hyponym",
	"#m": "Member holonym",
	"#s": "Substance holonym",
	"#p": "Part holonym",
	"%m": "Member meronym",
	"%s": "Substance meronym",
	"%p": "Part meronym",
	"-c": "Member of this domain - TOPIC",
	"-r": "Member of this domain - REGION",
	"-u": "Member of this domain - USAGE",

	// pointer_symbol
	"*": "Entailment",
	">": "Cause",
	"$": "Verb Group",

	// adjectives
	"&": "Similar to",
	"<": "Participle of verb",
	//"\\": "Pertainym (pertains to noun)",

	// adverbs
	//"\\": "Derived from adjective",
}

// https://wordnet.princeton.edu/documentation/wndb5wn
var synsetTypeMapping = map[string]string{
	"n": "noun",
	"v": "verb",
	"a": "adjective",
	"s": "adjective satellite",
	"r": "adverb",
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	synsets := parse.Synsets()

	for _, synset := range synsets[:20] {
		byt, _ := json.MarshalIndent(synset, "", "  ")
		log.Printf("%s", byt)
	}
}
