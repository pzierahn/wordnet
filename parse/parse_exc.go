package parse

import (
	_ "embed"
)

func Exc() (mapping map[string]string) {

	mapping = make(map[string]string)

	lines := readAll("", ".exc")
	parsedLines := parseLines(lines)

	for _, parts := range parsedLines {
		word := parts[len(parts)-1]

		for _, syn := range parts[:len(parts)-2] {
			mapping[syn] = word
		}
	}

	return
}
