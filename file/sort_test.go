package file

import (
	"io/ioutil"
	"os"
	"testing"
)

const FILE = "testdata/sort.txt"

func TestSort(t *testing.T) {
	_ = ioutil.WriteFile(FILE, SortMap(ReadFile(FILE)), os.ModeAppend)
}
