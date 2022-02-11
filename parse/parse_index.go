package parse

import (
	"log"
	"strconv"
)

type IndexFormat struct {
	Lemma         string // lemma
	Pos           string // pos
	SynsetCount   int    // synset_cnt
	SynsetPointer []int  // synset_offset
}

func Index() (records []IndexFormat) {

	// https://wordnet.princeton.edu/documentation/wndb5wn@
	// lemma  pos  synset_cnt  p_cnt  [ptr_symbol...]  sense_cnt  tagsense_cnt   synset_offset  [synset_offset...]

	lines := readAll("index", "", "index.sense")
	parsedLines := parseLines(lines)

	for _, line := range parsedLines {

		count, err := strconv.Atoi(line[2])
		if err != nil {
			log.Fatalf("cloudn't parse SynsetCount: %v", line)
		}

		var pointers []int
		for inx, str := range line[len(line)-count:] {
			pointer, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("cloudn't parse pointer(%d): %v", inx, line)
			}

			pointers = append(pointers, pointer)
		}

		entry := IndexFormat{
			Lemma:         line[0],
			Pos:           line[1],
			SynsetCount:   count,
			SynsetPointer: pointers,
		}

		records = append(records, entry)
	}

	return
}
