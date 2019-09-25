package file

import (
	"sync"
)

const LogFile = "./testdata/hello"

type File struct {
	mu   sync.RWMutex
	File string
}

func (f *File) SafeRead() {
	f.mu.RLock()
	defer f.mu.RUnlock()

}

func SafeWrite() {

}

func CreateFile() {

}

func OpenFile() {

}