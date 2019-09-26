package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"sync"
)

/* Read, Write, Append 都使用string进行交互*/
type File struct {
	mu   sync.RWMutex
	File string
}

const (
	WOnly  = 0
	Append = 1
)

type Operate interface {
	SafeRead() (string,error)
	SafeWrite(string) error      // overwrite
	SafeAppend(string) error     // append at end
}

// 增删改查
type Line interface {
	ReadLine(int) (string,error)
	WriteLine(string,int) error      // overwrite line
	AppendLine(string,int) error     // append end of line
	FindLine(string) int             // find in line
}

// 多线程安全读, 出参为string
func (f *File) SafeRead() (string,error){
	f.mu.RLock()
	defer f.mu.RUnlock()

	str, err := ioutil.ReadFile(f.File)
	return string(str), err
}

// 多线程安全写, 会覆盖整个文件
func (f *File) SafeWrite(data string) error{
	f.mu.Lock()
	defer f.mu.Unlock()

	return ioutil.WriteFile(f.File, []byte(data), os.ModeAppend)
}

// 多线程安全追加, 在文件末尾写
func (f *File) SafeAppend(data string) error{
	f.mu.Lock()
	defer f.mu.Unlock()

	old, err := ioutil.ReadFile(f.File)
	if err!=nil {
		old = []byte("")      // ignore errors
	}

	return ioutil.WriteFile(f.File, []byte(string(old) + data), os.ModeAppend)
}

// 读特定行，该行不存在则返回 ""
func (f *File) ReadLine(line int) (string,error){
	f.mu.RLock()
	defer f.mu.RUnlock()

	str, err := ioutil.ReadFile(f.File)
	sls := strings.Split(string(str), "\n")

	if len(sls)<line || line<=0{
		return "", err
	}
	return sls[line-1], err
}

// 写特定行, 文件无此行则创建此行
// 如 line=0, 则在最后一行写入
func (f *File) WriteLine(data string, line int) error{
	f.mu.Lock()
	defer f.mu.Unlock()

	old, err := ioutil.ReadFile(f.File)
	if err!=nil {
		old = []byte("")      // ignore errors
	}

	data = fixString(strings.Split(string(old),"\n"), data, line, WOnly)
	return ioutil.WriteFile(f.File, []byte(data), os.ModeAppend)
}

// 在特定行后追加, 文件无此行则创建此行
func (f *File) AppendLine(data string, line int) error{
	f.mu.Lock()
	defer f.mu.Unlock()

	old, err := ioutil.ReadFile(f.File)
	if err!=nil {
		old = []byte("")      // ignore errors
	}

	data = fixString(strings.Split(string(old),"\n"), data, line, Append)
	return ioutil.WriteFile(f.File, []byte(data), os.ModeAppend)
}

// 使用正则表达式匹配, 返回找到的第一行
func (f *File) FindLine(key string) int{
	f.mu.RLock()
	defer f.mu.RUnlock()

	str, _ := ioutil.ReadFile(f.File)
	sls := strings.Split(string(str), "\n")

	exp := regexp.MustCompile(key)
	for k := range sls {
		if exp.FindAllStringSubmatch(sls[k], -1) != nil {
			return k+1
		}
	}
	return -1
}

func fixString(sls []string, data string, line int, mode int) (res string){
	if line==0 {
		line = len(sls) +1
	}

	for i:=len(sls); i<line; i++ {
		sls = append(sls, "")
	}

	if line>0 {
		if mode==WOnly {
			sls[line-1] = data
		}else {
			sls[line-1] += data
		}
	}

	for _, v := range sls {
		res += v + "\n"
	}
	return
}

// example for file
func Example() {
	var testFile = File {
		File:    "./testdata/1.log",
	}
	_ = testFile.AppendLine("this is some append", 0)

	fmt.Println(testFile.ReadLine(20))
}