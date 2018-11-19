package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
	"runtime"
)

var i, j int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	fmt.Println("Hello, Go")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(1, 3))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(swap2("Hello", "World"))
	fmt.Println("i=", i, "j=", j)
	k := 4
	fmt.Println("k=", k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	var l int32 = 1
	m := int64(l)
	fmt.Println(m)
	v := 1
	fmt.Printf("%T\n", v)
	v = 3
	fmt.Printf("%T\n", v)
	const c = 1
	fmt.Printf("%v\n", c)
	fmt.Printf("%v", Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
	fmt.Println(sum())
	fmt.Println(sum2())
	fmt.Println(sum3())
	if_test(1)
	if_test(0)
	os()
	call_defer()
}

func call_defer() {
	for ; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func os() {
	var i = runtime.GOOS
	switch i {
	case "darwin":
		fmt.Println("Mac")
	default:
		fmt.Println(i)
	}
}

func if_test(i int) {
	if i > 0 {
		fmt.Println("Hello")
	} else {
		fmt.Println("World")
	}
}

func sum3() int {
	sum := 0
	for sum < 45 {
		sum += 1
	}
	return sum
}

func sum2() int {
	sum := 0
	for ; sum < 45; {
		sum += 1
	}
	return sum
}

func sum() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap2(x, y string) (a, b string) {
	a = y
	b = x
	return
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
