package datasort

import "sort"

var a = []int{10, 1, 9, 2, 8, 3, 7, 4, 6, 5}
var b = []float64{10, 1, 9, 2, 8, 3, 7, 4, 6, 5}
var c = []string{"test", "test"}

//接口方式实现
func shengxu() {
	//Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。
	//本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）。
	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.Float64Slice(b))
	sort.Sort(sort.StringSlice(c))
	//Stable排序data，并保证排序的稳定性，相等元素的相对次序不变。
	//它调用1次data.Len，O(n*log(n))次data.Less和O(n*log(n)*log(n))次data.Swap。
	sort.Stable(sort.IntSlice(a))
	sort.Stable(sort.Float64Slice(b))
	sort.Stable(sort.StringSlice(c))
	//IsSorted报告data是否已经被排序。
	intbool := sort.IsSorted(sort.IntSlice(a))
	intboo2 := sort.IsSorted(sort.Float64Slice(b))
	intboo3 := sort.IsSorted(sort.StringSlice(c))
	print(intbool, intboo2, intboo3)
}

//接口方式实现
func jiangxu() {
	//Reverse包装一个Interface接口并返回一个新的Interface接口，对该接口排序可生成递减序列。
	x := sort.Reverse(sort.IntSlice(a))
	y := sort.Reverse(sort.Float64Slice(b))
	z := sort.Reverse(sort.StringSlice(c))
	//Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。
	//本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）。
	sort.Sort(x)
	sort.Sort(y)
	sort.Sort(z)
	//Stable排序data，并保证排序的稳定性，相等元素的相对次序不变。
	//它调用1次data.Len，O(n*log(n))次data.Less和O(n*log(n)*log(n))次data.Swap。
	sort.Stable(x)
	sort.Stable(y)
	sort.Stable(z)
}

//函数方式实现
func sheng() {
	//排序为递增顺序。
	sort.Ints(a)
	sort.Float64s(b)
	sort.Strings(c)
	//检查a是否已排序为递增顺序。
	sort.IntsAreSorted(a)
	sort.Float64sAreSorted(b)
	sort.StringsAreSorted(c)
	//在递增顺序的a中搜索x，返回x的索引。如果查找不到，
	//返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	int1 := sort.SearchInts(a, 5)
	int2 := sort.SearchFloat64s(b, 5)
	int3 := sort.SearchStrings(c, "test")
	print(int1, int2, int3)
}

//函数方式实现
func jiang() {
	//noway
}
