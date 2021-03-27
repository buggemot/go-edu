package csvfile

import (
	"os"
	"encoding/csv"
	"log"
	"strings"
)

type csvFile struct {
	name string
	body []byte
	records [][]string
}

func (cf *csvFile) Read() {
	var err error
	cf.body, err = os.ReadFile(cf.name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	r := csv.NewReader(strings.NewReader(string(cf.body)))
	cf.records, err = r.ReadAll()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
