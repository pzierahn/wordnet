package parse

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

const wordnetDBPath = "input"

func readFileLines(files ...string) (lines []string) {
	for _, filename := range files {
		bys, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("couldn't read file %s", filename)
		}

		content := string(bys)
		parts := strings.Split(content, "\n")

		for _, line := range parts {
			// Comment line, legal stuff, etc.
			if strings.HasPrefix(line, " ") {
				continue
			}

			line = strings.TrimSpace(line)
			if line != "" {
				lines = append(lines, line)
			}
		}
	}

	return
}

func readAll(prefix, suffix string, exceptions ...string) (lines []string) {
	all, err := ioutil.ReadDir(wordnetDBPath)
	if err != nil {
		log.Fatalf("cloudn't read dir: %v", err)
	}

	var files []string

loop:
	for _, file := range all {

		if !strings.HasPrefix(file.Name(), prefix) {
			continue
		}

		if !strings.HasSuffix(file.Name(), suffix) {
			continue
		}

		for _, exception := range exceptions {
			if file.Name() == exception {
				continue loop
			}
		}

		filename := filepath.Join(wordnetDBPath, file.Name())
		files = append(files, filename)
	}

	return readFileLines(files...)
}

func parseLines(lines []string) (records [][]string) {
	for _, line := range lines {
		records = append(records, strings.Split(line, " "))
	}

	return
}
