package csvfile

import (
	"testing"
	"strings"
)

func TestRead(t *testing.T) {
	cf := &csvFile{name: "pattern.csv"}
	cf.Read()
	if strings.Trim(cf.records[1][2], " ") != "This is also a test string" {
		t.Fatalf("Incorrect value: check position[1][2] in array\n")
	}

}