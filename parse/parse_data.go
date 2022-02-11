package parse

import (
	"log"
	"strconv"
	"strings"
)

type Pointer struct {
	Symbol string // pointer_symbol
	Id     int    // synset_offset
	Pos    string // pos
	Source int    // source
	Target int    // target
}

type Word struct {
	Word  string // word
	LexId int    // lex_id
}

type Synset struct {
	Id                 int       // synset_offset
	SynsetType         string    // ss_type
	WordCount          int       // w_cnt
	Words              []Word    // word + lex_id
	SynsetPointerCount int       // p_cnt
	SynsetPointers     []Pointer // ptr
	Definitions        []string  // gloss
	Examples           []string  // gloss
}

func Synsets() (synsets []Synset) {
	// https://wordnet.princeton.edu/documentation/wndb5wn

	data := readAll("data.", "")

	for _, record := range data {
		id, err := strconv.Atoi(record[:8])
		if err != nil {
			log.Fatalf("couldn't parse record '%v'", record)
		}

		parts := strings.Split(record, " | ")

		meta := parts[0]

		metaParts := strings.Split(meta, " ")

		var definitions []string
		var examples []string

		gloss := strings.Split(parts[1], "; ")
		for _, fragment := range gloss {
			//examples[iny] = strings.Trim(examples[iny], " \"'")
			if strings.HasPrefix(fragment, "\"") {
				examples = append(examples, strings.Trim(fragment, " \"'"))
			} else {
				definitions = append(definitions, strings.TrimSpace(fragment))
			}
		}

		// n    NOUN
		// v    VERB
		// a    ADJECTIVE
		// s    ADJECTIVE SATELLITE
		// r    ADVERB
		ssType := metaParts[2]

		wCnt64, err := strconv.ParseInt(metaParts[3], 16, 64)
		if err != nil {
			log.Fatalf("couldn't parse wCnt '%v' in %v", metaParts[3], metaParts)
		}

		wCnt := int(wCnt64)

		offset := 4

		var words []Word
		for iny := 0; iny < wCnt; iny++ {
			word := metaParts[offset+iny*2]
			lexId := metaParts[offset+iny*2+1]

			lexId64, err := strconv.ParseInt(lexId, 16, 64)
			if err != nil {
				log.Fatalf("cloudn't parse int %v", metaParts)
			}

			words = append(words, Word{
				Word:  word,
				LexId: int(lexId64),
			})
		}

		offset += wCnt * 2

		pCnt, err := strconv.Atoi(metaParts[offset])
		if err != nil {
			log.Fatalf("couldn't parse pCnt '%v'", metaParts[offset])
		}

		offset += 1

		var pointers []Pointer

		// pointer_symbol  synset_offset  pos  source/target
		for iny := 0; iny < pCnt; iny++ {
			start := offset + iny*4
			end := start + 4

			ptr := metaParts[start:end]

			id, err := strconv.Atoi(ptr[1])
			if err != nil {
				log.Fatalf("couldn't parse ptr synset_offset %v", ptr)
			}

			source, err := strconv.ParseInt(ptr[3][:2], 16, 64)
			target, err := strconv.ParseInt(ptr[3][2:], 16, 64)

			pointer := Pointer{
				Symbol: ptr[0],
				Id:     id,
				Pos:    ptr[2],
				Source: int(source),
				Target: int(target),
			}

			pointers = append(pointers, pointer)
		}

		synset := Synset{
			Id:                 id,
			SynsetType:         ssType,
			WordCount:          wCnt,
			Words:              words,
			SynsetPointerCount: pCnt,
			SynsetPointers:     pointers,
			Definitions:        definitions,
			Examples:           examples,
		}

		//log.Printf("id=%08d", id)
		//log.Printf("meta=%v", meta)
		//log.Printf("ssType=%v (pos)", ssType)
		//log.Printf("wCnt=%v", wCnt)
		//log.Printf("words=%v", words)
		//log.Printf("pCnt=%v", pCnt)
		//log.Printf("ptr=%v", pointers)
		//log.Printf("definition=%s", definition)

		//byt, _ := json.MarshalIndent(sense, "", "  ")
		//log.Printf("sense=%s", byt)

		synsets = append(synsets, synset)
	}

	return
}
