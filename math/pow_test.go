package math

import (
	"fmt"
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	fmt.Println(MyPow(1.2, 21))
	fmt.Println(math.Pow(1.2,21))
}