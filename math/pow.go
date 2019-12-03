package math

func MyPow(x float64, y int64) float64 {
	// ans = a1 * 2**ae (= 1 for now).
	a1, ae := 1.0, 0

	x1, xe := Frexp(x)
	for i := y; i != 0; i >>= 1 {

		// fmt.Printf("i=%2d, a1=%8f, ae=%2d, x1=%8f, xe=%2d\n", i, a1, ae, x1, xe)

		if i&1 == 1 {   // 奇数  <--- zb
			a1 *= x1
			ae += xe
		}

		x1 *= x1
		xe <<= 1

		if x1 < .5 {    // 防止 a1 < 0.1
			x1 += x1
			xe--
		}
	}

	// fmt.Println("a1=", a1, " ae=", ae)
	return Ldexp(a1, ae)
}

// 将 x 分解为 x1 * 2^xe
func Frexp(x float64) (float64, int) {
	var xe = 0

	for ; x>=1; x/=2 {
		xe++
	}
	return x, xe
}

// 将 x1 * 2^xe 合并为 x
func Ldexp(x1 float64, xe int) float64 {
	return x1 * float64(int(1) << xe)
}
