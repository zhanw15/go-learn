package change

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// 10进制到2进制
func DecBin(n int64) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

// 10进制到8进制
func DecOct(d int64) int64 {
	if d == 0 {
		return 0
	}
	if d < 0 {
		log.Println("Decimal to octal error: the argument must be greater than zero.")
		return -1
	}
	s := ""
	for q := d; q > 0; q = q / 8 {
		m := q % 8
		s = fmt.Sprintf("%v%v", m, s)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Println("Decimal to octal error:", err.Error())
		return -1
	}
	return int64(n)
}

// 10进制到16进制
func DecHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

// 二进制到十进制
func BinDec(b string) (n int64) {
	s := strings.Split(b, "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Binary to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int64(d)
}

// 八进制到十进制
func OctDec(o int64) (n int64) {
	s := strings.Split(strconv.Itoa(int(o)), "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Octal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(8, float64(l-i-1))
	}
	return int64(d)
}

// 16进制到10进制
func HexDec(h string) (n int64) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			log.Println("Hexadecimal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}

// 八进制到2进制
func OctBin(o int64) string {
	d := OctDec(o)
	if d == -1 {
		return ""
	}
	return DecBin(d)
}

// 16进制到2进制
func HexBin(h string) string {
	d := HexDec(h)
	if d == -1 {
		return ""
	}
	return DecBin(d)
}

// 2进制到8进制
func BinOct(b string) int64 {
	d := BinDec(b)
	if d == -1 {
		return -1
	}
	return DecOct(d)
}

// 2进制到16进制
func BinHex(b string) string {
	d := BinDec(b)
	if d == -1 {
		return ""
	}
	return DecHex(d)
}
