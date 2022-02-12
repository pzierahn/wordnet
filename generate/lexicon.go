package generate

import (
	"wordnet/parse"
	"wordnet/utils"
)

type LemmaPointers map[string][]int
type SynsetIndex map[int]parse.Synset

func Lexicon() {

	index := parse.Index()
	lemmaPointers := make(LemmaPointers)
	for _, entry := range index {
		//log.Printf("Word=%s pointers=%v", entry.Lemma, entry.SynsetPointer)
		word := entry.Lemma
		lemmaPointers[word] = append(lemmaPointers[word], entry.SynsetPointer...)
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

	exceptions := parse.Exc()
	utils.ExportJson("lexicon/exceptions.json", exceptions)
	utils.ExportGob("lexicon/exceptions.gob", exceptions)

	return
}
