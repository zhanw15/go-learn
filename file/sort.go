package file

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func SortMap(I []int, S []string) []byte {
	sort.Ints(I)
	sort.Strings(S)

	var str = ""
	for _, v := range I {
		str += strconv.Itoa(v) + "\n"
	}
	for _, v := range S {
		str += v + "\n"
	}

	return []byte(str)
}

func ReadFile(file string) (I []int, S []string){
	str, err := ioutil.ReadFile(file)
	if err!=nil {
		return nil, nil
	}
	sls := strings.FieldsFunc(string(str), split)

	fmt.Println(sls)

	for _, v := range sls {
		if i, ok := strconv.Atoi(v); ok!=nil {
			S = append(S, v)
		}else {
			I = append(I, i)
		}
	}
	return
}

// 多分隔符排序
func split(r rune) bool {
	return r=='\n' || r==' '
}