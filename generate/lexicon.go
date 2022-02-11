package generate

import (
	"wordnet/parse"
	"wordnet/utils"
)

type LemmaPointers map[string]map[int]bool
type SynsetIndex map[int]parse.Synset

func Lexicon() {

	index := parse.Index()
	lemmaPointers := make(LemmaPointers)
	for _, entry := range index {
		//log.Printf("Word=%s pointers=%v", entry.Lemma, entry.SynsetPointer)
		word := entry.Lemma

		if _, exist := lemmaPointers[word]; !exist {
			lemmaPointers[word] = make(map[int]bool)
		}

		for _, pointer := range entry.SynsetPointer {
			lemmaPointers[word][pointer] = true
		}
	}

	utils.ExportJson("lexicon/lemmaPointers.json", lemmaPointers)
	utils.ExportGob("lexicon/lemmaPointers.gob", lemmaPointers)

	synsets := parse.Synsets()
	synsetIndex := make(SynsetIndex)
	for _, synset := range synsets {
		synsetIndex[synset.Id] = synset
	}

	utils.ExportJson("lexicon/synsetIndex.json", synsetIndex)
	utils.ExportGob("lexicon/synsetIndex.gob", synsetIndex)

	return
}
