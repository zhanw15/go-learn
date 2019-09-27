package example

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Worker struct {
	Reporter *sync.WaitGroup
}

type data struct {
	Info [7]string
}

func ExampleFile()  {
	start := time.Now()
	fmt.Println("work started!")

	worker := NewWorker()

	for i:=0; i<2; i++ {
		worker.Reporter.Add(1)
		go worker.work(i%2==0)
	}

	worker.Reporter.Wait()

	fmt.Printf("all work finished! spend time: %v\n", time.Now().Sub(start))
}

func NewWorker() Worker {
	return Worker{Reporter: new(sync.WaitGroup)}
}

func (w *Worker) work(b bool) {
	defer w.Reporter.Done()

	str, err := ioutil.ReadFile("file/testdata/data-m2.txt")
	if err!=nil {
		panic("Worker Done! : can't read file!")
	}

	sls := strings.Split(string(str), "\n")

	var Data = make([]data, len(sls))
	for k, v := range sls {
		sub := strings.Split(v, "|")
		for i := range sub {
			Data[k].Info[i] = sub[i]
		}
	}

	sort.Sort(DataType(Data))

	var str2 = ""

	if b {
		for k := range Data {
			if !ouShu(Data[k].Info[0]) {
				for _, v := range Data[k].Info {
					if v!="" {
						str2  += v + "|"
					}
				}
				if Data[k].Info[0]!="" {
					str2 = str2[:len(str2)-1]
					str2 += "\n"
				}
			}
		}
		err := ioutil.WriteFile("file/testdata/A.txt", []byte(str2), os.ModeAppend)
		if err!=nil {
			panic("Worker Done! : can't write file A!")
		}
	}else {
		for k := range Data {
			if ouShu(Data[k].Info[0]) {
				for _, v := range Data[k].Info {
					if v!="" {
						str2  += v + "|"
					}
				}
				if Data[k].Info[0]!="" {
					str2 = str2[:len(str2)-1]
					str2 += "\n"
				}
			}
		}
		err := ioutil.WriteFile("file/testdata/B.txt", []byte(str2), os.ModeAppend)
		if err!=nil {
			panic("Worker Done! : can't write file B!")
		}
	}
}

// 排序接口
type DataType []data

func (a DataType) Len() int {
	return len(a)
}

func (a DataType) Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}

func (a DataType) Less(i, j int) bool {

	if a[i].Info[0] == a[j].Info[0] {
		i1, _ := strconv.Atoi(a[i].Info[1])
		i2, _ := strconv.Atoi(a[j].Info[1])
		return i1 < i2
	}
	return a[i].Info[0] < a[j].Info[0]
}

func ouShu(s string) bool{

	if s!="" {
		s = s[len(s)-1:]
	}
	if n, ok := strconv.Atoi(s); ok==nil {
		return n%2==0
	}
	switch s {
	case "a", "c", "e": return true
	}
	return false
}