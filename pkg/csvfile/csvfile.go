package csvfile

import (
	"os"
	"encoding/csv"
	"log"
	"strings"
)

type CsvFile struct {
	Name string
	Body []byte
	Records [][]string
}

func NewCsvFile() *CsvFile {
	return &CsvFile{}
}

func (cf *CsvFile) Read() {
	var err error
	cf.Body, err = os.ReadFile(cf.Name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	r := csv.NewReader(strings.NewReader(string(cf.Body)))
	cf.Records, err = r.ReadAll()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
