package lexicon

import (
	_ "embed"
	"log"
	"time"
	"wordnet/generate"
	"wordnet/utils"
)

var (
	//go:embed lemmaPointers.gob
	lemmasRaw []byte
	//go:embed synsetIndex.gob
	synsetsRaw []byte
	//go:embed exceptions.gob
	exceptionsRaw []byte

	index      generate.LemmaPointers
	data       generate.SynsetIndex
	exceptions map[string]string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	start := time.Now()

	if err := utils.UnmarshalGob(lemmasRaw, &index); err != nil {
		log.Fatalf("error UnmarshalGob: %v", err)
	}

	if err := utils.UnmarshalGob(synsetsRaw, &data); err != nil {
		log.Fatalf("error UnmarshalGob: %v", err)
	}

	if err := utils.UnmarshalGob(exceptionsRaw, &exceptions); err != nil {
		log.Fatalf("error UnmarshalGob: %v", err)
	}

	log.Printf("init time: %v", time.Now().Sub(start))
	//log.Printf("index=%v", len(index))
	//log.Printf("data=%v", len(data))
}
