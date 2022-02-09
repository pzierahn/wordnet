package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"wordnet/parse"
	"wordnet/utils"
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

	// verbs
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

type Sense struct {
	Index       int
	Pos         string
	Synonyms    []string
	Definitions []string
	Examples    []string
	Misc        map[string][]string
}

type LexiconEntry struct {
	Word   string
	Senses []Sense
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	synsets := parse.Synsets()

	wordIndex := make(map[int][]string)

	for _, synset := range synsets {
		for _, word := range synset.Words {
			wordIndex[synset.Id] = append(wordIndex[synset.Id], word.Word)
		}
	}

	lex := LexiconEntry{
		Word: "count",
	}

	for _, synset := range synsets[:] {

		var contains bool

		for _, word := range synset.Words {
			if word.Word == lex.Word {
				contains = true
				break
			}
		}

		if !contains {
			continue
		}

		sense := Sense{
			Index:       0,
			Pos:         synsetTypeMapping[synset.SynsetType],
			Synonyms:    nil,
			Definitions: synset.Definitions,
			Examples:    synset.Examples,
			Misc:        make(map[string][]string),
		}

		for _, word := range synset.Words {
			if word.Word != lex.Word {
				sense.Synonyms = append(sense.Synonyms, word.Word)
			} else {
				sense.Index = word.LexId
			}
		}

		for _, pointer := range synset.SynsetPointers {
			name := symbolMapping[pointer.Symbol]
			list := wordIndex[pointer.Id]
			sense.Misc[name] = append(sense.Misc[name], list...)
		}

		for name, list := range sense.Misc {
			sense.Misc[name] = utils.RemoveDuplicateStrings(list)
			sort.Strings(sense.Misc[name])
		}

		lex.Senses = append(lex.Senses, sense)

		//byt, _ := json.MarshalIndent(synset, "", "  ")
		//log.Printf("%s", byt)
	}

	sort.Slice(lex.Senses, func(i, j int) bool {
		return lex.Senses[i].Index < lex.Senses[j].Index
	})

	byt, _ := json.MarshalIndent(lex, "", "  ")
	log.Printf("lex=%s", byt)

	_ = ioutil.WriteFile("dumb/dumb.json", byt, 0755)
}
