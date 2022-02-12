package lexicon

import (
	"github.com/pzierahn/wordnet/parse"
	"github.com/pzierahn/wordnet/utils"
	"sort"
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
	Pos         string
	Synonyms    []string
	Definitions []string
	Examples    []string
	Misc        map[string][]string
}

type Entry struct {
	Word   string
	Senses []Sense
}

// Get returns all definitions.
func Get(search string) (entry Entry) {

	if word, ok := exceptions[search]; ok {
		search = word
	}

	entry = Entry{
		Word: search,
	}

	var synsets []parse.Synset
	for _, pointer := range index[search] {
		synsets = append(synsets, data[pointer])
	}

	for _, synset := range synsets {
		sense := Sense{
			Pos:         synsetTypeMapping[synset.SynsetType],
			Definitions: synset.Definitions,
			Examples:    synset.Examples,
			Misc:        make(map[string][]string),
		}

		for _, word := range synset.Words {
			if word.Word != entry.Word {
				sense.Synonyms = append(sense.Synonyms, word.Word)
			}
		}

		for _, pointer := range synset.SynsetPointers {
			name := symbolMapping[pointer.Symbol]

			for _, word := range data[pointer.Id].Words {
				sense.Misc[name] = append(sense.Misc[name], word.Word)
			}
		}

		for name, list := range sense.Misc {
			sense.Misc[name] = utils.RemoveDuplicateStrings(list)
			sort.Strings(sense.Misc[name])
		}

		entry.Senses = append(entry.Senses, sense)

		//byt, _ := json.MarshalIndent(synset, "", "  ")
		//log.Printf("%s", byt)
	}

	return
}

// Words returns all words in the lexicon.
func Words() (words []string) {
	for word := range index {
		words = append(words, word)
	}

	sort.Strings(words)

	return
}
