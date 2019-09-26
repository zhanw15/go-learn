package file

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var(
	NoKEY   = errors.New("NO SUCK KEY FIND")
//	BadFile = errors.New("Wrong Disk File ")

	etcd    = File{File:"./testdata/etc.d"}
)

// 属性增删改查, 线程不安全
type Operation interface {
	Set(string,string)   error
	Get(string)          (string,error)
	Delete(string)       error
}

func (f *File) Set(key string, value string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	m := getMap(f.File)

	m[key] = value
	data, _ := json.Marshal(m)
	return ioutil.WriteFile(f.File, data, os.ModeAppend)
}

func (f *File) Get(key string) (string, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	m := getMap(f.File)

	if value, ok := m[key]; ok {
		return value, nil
	}
	return "", NoKEY
}

func (f *File) Delete(key string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	m := getMap(f.File)

	// can't find key
	if _, ok := m[key]; !ok {
		return NoKEY
	}

	// find key, delete
	delete(m, key)
	data, _ := json.Marshal(m)
	return ioutil.WriteFile(f.File, data, os.ModeAppend)
}

func getMap(f string) map[string]string{
	m := make(map[string]string, 0)
	str, err := ioutil.ReadFile(f)   // ignore errors
	if err!=nil {
		str = []byte("{}")
	}
	_ = json.Unmarshal(str, &m)

	return m
}