package json

import (
	"encoding/json"
	"fmt"
)

func Firstdecod() {
	var book Book
	gobook := Book{
		"Go语言编程",
		[]string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99}
	b, _ := json.Marshal(gobook)
	json.Unmarshal(b, &book)
	fmt.Printf("gobook:%v\n", book)
	fmt.Printf("gobook:%+v\n", book)
}

func Seconddecod() {
	var ebook Book
	var book Book
	gobook := &Book{
		"Go语言编程",
		[]string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99}
	emptybook := &Book{}
	a, _ := json.Marshal(emptybook)
	b, _ := json.Marshal(gobook)
	json.Unmarshal(a, &ebook)
	json.Unmarshal(b, &book)
	fmt.Printf("emptybook:%+v\n,gobook:%+v\n", ebook, book)
}

func Thirddecod() {
	type Struct struct {
		A int
		b int
	}
	var aa bool
	var bb int
	var cc float64
	var dd string
	var ee [4]interface{}
	var ff []interface{}
	var gg Struct
	hh := make(map[string][]interface{})
	ii := make(map[int]int)
	jj := make(map[float64]float64)
	kk := make(map[bool]bool)
	Bool := true
	Int := 1
	Float64 := 1.1
	String := "<string>"
	Array := [4]interface{}{Bool, Int, Float64, String}
	Slice := []interface{}{Bool, Int, Float64, String}
	S := Struct{Int, Int}
	Map := make(map[string][]interface{})
	Map["Array"] = Slice
	Mapint := make(map[int]int)
	Mapint[Int] = Int
	Mapfloat64 := make(map[float64]float64)
	Mapfloat64[Float64] = Float64
	Mapbool := make(map[bool]bool)
	Mapbool[Bool] = Bool
	//Mapslice := make(map[[]interface{}][]interface{})
	//Mapstruct := make(map[struct]struct)
	a, _ := json.Marshal(Bool)
	b, _ := json.Marshal(Int)
	c, _ := json.Marshal(Float64)
	d, _ := json.Marshal(String)
	e, _ := json.Marshal(Array)
	f, _ := json.Marshal(Slice)
	g, _ := json.Marshal(S)
	h, _ := json.Marshal(Map)
	i, _ := json.Marshal(Mapint)
	j, _ := json.Marshal(Mapfloat64)
	k, _ := json.Marshal(Mapbool)
	json.Unmarshal(a, &aa)
	json.Unmarshal(b, &bb)
	json.Unmarshal(c, &cc)
	json.Unmarshal(d, &dd)
	json.Unmarshal(e, &ee)
	json.Unmarshal(f, &ff)
	json.Unmarshal(g, &gg)
	json.Unmarshal(h, &hh)
	json.Unmarshal(i, &ii)
	json.Unmarshal(j, &jj)
	json.Unmarshal(k, &kk)
	fmt.Printf("a: %T,%v\n", aa, aa)
	fmt.Printf("b: %T,%v\n", bb, bb)
	fmt.Printf("c: %T,%v\n", cc, cc)
	fmt.Printf("d: %T,%v\n", dd, dd)
	fmt.Printf("e: %T,%v\n", ee, ee)
	fmt.Printf("f: %T,%v\n", ff, ff)
	fmt.Printf("g: %T,%v\n", gg, gg)
	fmt.Printf("h: %T,%v\n", hh, hh)
	fmt.Printf("i: %T,%v\n", ii, ii)
	fmt.Printf("j: %T,%v\n", jj, jj)
	fmt.Printf("k: %T,%v\n", kk, kk)
}

func Fourthdecod() {
	var book Book
	//b := Book{Title: "Go语言编程", Sales: 1000000}
	b := []byte(`{"Title": "Go语言编程", "Sales": 1000000}`)
	err := json.Unmarshal(b, &book)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("book: %T,%v\n", book, book)
}

func Fifthdecod() {
	var aa interface{}
	var bb interface{}
	var cc interface{}
	var dd interface{}
	var ee interface{}
	var ff interface{}
	var gg interface{}
	var hh interface{}
	var ii interface{}
	var jj interface{}
	var kk interface{}
	type Struct struct {
		A int
		b int
	}
	Bool := true
	Int := 1
	Float64 := 1.1
	String := "<string>"
	Array := [4]interface{}{Bool, Int, Float64, String}
	Slice := []interface{}{Bool, Int, Float64, String}
	S := Struct{Int, Int}
	Map := make(map[string][]interface{})
	Map["Array"] = Slice
	Mapint := make(map[int]int)
	Mapint[Int] = Int
	Mapfloat64 := make(map[float64]float64)
	Mapfloat64[Float64] = Float64
	Mapbool := make(map[bool]bool)
	Mapbool[Bool] = Bool
	//Mapslice := make(map[[]interface{}][]interface{})
	//Mapstruct := make(map[struct]struct)
	a, _ := json.Marshal(Bool)
	b, _ := json.Marshal(Int)
	c, _ := json.Marshal(Float64)
	d, _ := json.Marshal(String)
	e, _ := json.Marshal(Array)
	f, _ := json.Marshal(Slice)
	g, _ := json.Marshal(S)
	h, _ := json.Marshal(Map)
	i, _ := json.Marshal(Mapint)
	j, _ := json.Marshal(Mapfloat64)
	k, _ := json.Marshal(Mapbool)
	json.Unmarshal(a, &aa)
	json.Unmarshal(b, &bb)
	json.Unmarshal(c, &cc)
	json.Unmarshal(d, &dd)
	json.Unmarshal(e, &ee)
	json.Unmarshal(f, &ff)
	json.Unmarshal(g, &gg)
	json.Unmarshal(h, &hh)
	json.Unmarshal(i, &ii)
	json.Unmarshal(j, &jj)
	json.Unmarshal(k, &kk)
	fmt.Printf("a: %T,%v\n", aa, aa)
	fmt.Printf("b: %T,%v\n", bb, bb)
	fmt.Printf("c: %T,%v\n", cc, cc)
	fmt.Printf("d: %T,%v\n", dd, dd)
	fmt.Printf("e: %T,%v\n", ee, ee)
	fmt.Printf("f: %T,%v\n", ff, ff)
	fmt.Printf("g: %T,%v\n", gg, gg)
	fmt.Printf("h: %T,%v\n", hh, hh)
	fmt.Printf("i: %T,%v\n", ii, ii)
	fmt.Printf("j: %T,%v\n", jj, jj)
	fmt.Printf("k: %T,%v\n", kk, kk)
}

func Sixthdecod() {
	gobook := Book{
		"Go语言编程",
		[]string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99,
	}
	b, _ := json.Marshal(gobook)
	var r interface{}
	json.Unmarshal(b, &r)
	book, ok := r.(map[string]interface{})
	if ok {
		for k, v := range book {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}
