package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type Cache interface {
	Add(k string, v string)
	Update(k string, v string)
	Remove(k string)
	Get(k string) string
	Clean()
	Println()
	Size() int
}

type DataCache struct {
	mu   sync.RWMutex
	File string
}

func NewStringCache() Cache{
	cache := new(DataCache)
	cache.File = time.Now().Format("2006-01-02-15-04-05-") + strconv.Itoa(rand.Intn(1000)) +".txt"
	os.Create(cache.File)

	return cache
}

func (f *DataCache) Update(k string, v string)  {
	f.mu.Lock()
	defer f.mu.Unlock()

	m := getMap(f.File)

	m[k] = v
	data, _ := json.Marshal(m)
	_ = ioutil.WriteFile(f.File, data, os.ModeAppend)
}

func (f *DataCache) Size() int {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return len(getMap(f.File))
}

func (f *DataCache) Println() {
	f.mu.Lock()
	defer f.mu.Unlock()

	m := getMap(f.File)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func (f *DataCache) Add(k string, v string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	m := getMap(f.File)

	m[k] = v
	data, _ := json.Marshal(m)
	err := ioutil.WriteFile(f.File, data, os.ModeAppend)

	if err!=nil {
		fmt.Println(err)
	}
}

func (f *DataCache) Get(k string) string {
	f.mu.RLock()
	defer f.mu.RUnlock()

	m := getMap(f.File)

	if value, ok := m[k]; ok {
		return value
	}
	return ""
}

func (f *DataCache) Clean()  {
	f.mu.Lock()
	defer f.mu.Unlock()

	ioutil.WriteFile(f.File, []byte("{}"), os.ModeAppend)
}

func (f *DataCache) Remove(k string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	m := getMap(f.File)

	// can't find key
	if _, ok := m[k]; !ok {
		return
	}

	// find key, delete
	delete(m, k)
	data, _ := json.Marshal(m)
	_ = ioutil.WriteFile(f.File, data, os.ModeAppend)
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

func main() {
	//开发字符串缓存程序，按缓存接口实现功能
	//要求：满足以下测试用例
	cache := NewStringCache()
	fmt.Println("add 3 items to cache ----")
	cache.Add("key1","value1")
	cache.Add("key2","value2")
	cache.Add("key3","value3")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("get key2 data from cache ----")
	fmt.Println(cache.Get("key2"))
	fmt.Println("add another 3 items to cache ----")
	cache.Add("key4","value4")
	cache.Add("key5","value5")
	cache.Add("key6","value6")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("update key5 to value88888 ----")
	cache.Update("key5","value88888")
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("update key9 to value99999 ----")
	cache.Update("key9","value99999")
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("remove key5 item from cache ----")
	fmt.Println(cache.Get("key5"))
	cache.Remove("key5")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("remove key7 item from cache ----")
	cache.Remove("key7")
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("clear the cache datas ----")
	cache.Clean()
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
	fmt.Println("add 100 datas to cache ----")
	for i:=0;i<100;i++ {
		go cache.Add("key"+strconv.Itoa(i),"value"+strconv.Itoa(i))
	}
	time.Sleep(15)
	fmt.Println("get cache size ----")
	fmt.Println(cache.Size())
	fmt.Println("print cache datas ----")
	cache.Println()
}