package json

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func Firstencod() {
	gobook := Book{
		"Go语言编程",
		[]string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99}
	b, _ := json.Marshal(gobook)
	fmt.Printf("gobook:%v\n", b)
	fmt.Printf("gobook:%v\n", string(b))
}

func Secondencod() {
	gobook := &Book{
		"Go语言编程",
		[]string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99}
	emptybook := &Book{}
	a, _ := json.Marshal(emptybook)
	b, _ := json.Marshal(gobook)
	fmt.Printf("emptybook:%v\n,gobook:%v\n", a, b)
	fmt.Printf("emptybook:%v\n,gobook:%v\n", string(a), string(b))
}

func Thirdencod() {
	Bool := true
	Int := 1
	Float64 := 1.1
	String := "<string>"
	Array := [4]interface{}{Bool, Int, Float64, String}
	Slice := []interface{}{Bool, Int, Float64, String}
	type Struct struct {
		A int
		b int
	}
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
	i, err := json.Marshal(Mapint)
	if err != nil {
		fmt.Print(err)
	}
	j, err := json.Marshal(Mapfloat64)
	if err != nil {
		fmt.Print(err)
	}
	k, err := json.Marshal(Mapbool)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("a: %T,%v\n", a, a)
	fmt.Printf("a: %T,%v\n", string(a), string(a))
	fmt.Printf("b: %T,%v\n", b, b)
	fmt.Printf("b: %T,%v\n", string(b), string(b))
	fmt.Printf("c: %T,%v\n", c, c)
	fmt.Printf("c: %T,%v\n", string(c), string(c))
	fmt.Printf("d: %T,%v\n", d, d)
	fmt.Printf("d: %T,%v\n", string(d), string(d))
	fmt.Printf("e: %T,%v\n", e, e)
	fmt.Printf("e: %T,%v\n", string(e), string(e))
	fmt.Printf("f: %T,%v\n", f, f)
	fmt.Printf("f: %T,%v\n", string(f), string(f))
	fmt.Printf("g: %T,%v\n", g, g)
	fmt.Printf("g: %T,%v\n", string(g), string(g))
	fmt.Printf("h: %T,%v\n", h, h)
	fmt.Printf("h: %T,%v\n", string(h), string(h))
	fmt.Printf("i: %T,%v\n", i, i)
	fmt.Printf("i: %T,%v\n", string(i), string(i))
	fmt.Printf("j: %T,%v\n", j, j)
	fmt.Printf("j: %T,%v\n", string(j), string(j))
	fmt.Printf("k: %T,%v\n", k, k)
	fmt.Printf("k: %T,%v\n", string(k), string(k))
}
