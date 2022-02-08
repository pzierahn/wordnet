package assets

import (
	"embed"
	_ "embed"
	"log"
	"strings"
)

//go:embed db
var fs embed.FS

func ReadFiles(files ...string) (records [][]string) {

	for _, filename := range files {
		bys, err := fs.ReadFile(filename)
		if err != nil {
			log.Fatalf("couldn't read %s", filename)
		}

		content := string(bys)
		lines := strings.Split(content, "\n")

		for _, line := range lines {
			if strings.HasPrefix(line, " ") {
				continue
			}

			line = strings.TrimSpace(line)
			record := strings.Split(line, " ")

			if len(record) > 1 {
				records = append(records, record)
			}
		}
	}

	return
}

func IndexPOS() (records [][]string) {
	return ReadFiles(
		"db/index.adj",
		"db/index.adv",
		"db/index.noun",
		"db/index.verb",
	)
}

func DataPOS() (records []string) {
	files := []string{
		"db/data.adj",
		"db/data.adv",
		"db/data.noun",
		"db/data.verb",
	}

	for _, filename := range files {
		bys, err := fs.ReadFile(filename)
		if err != nil {
			log.Fatalf("couldn't read %s", filename)
		}

		content := string(bys)
		lines := strings.Split(content, "\n")

		for _, line := range lines {
			if strings.HasPrefix(line, " ") {
				continue
			}

			line = strings.TrimSpace(line)

			if line != "" {
				records = append(records, line)
			}
		}
	}

	return

}
