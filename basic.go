package main

import (
    "fmt"
    "math"
    "math/rand"
    "math/cmplx"
)

// simple function
func add(x int, y int) int {
    return x + y
}

// multiple return
func swap(x, y string) (string, string) {
    return y, x
}

// named return
func split(sum int) (x, y int) {
    x = sum % 10
    y = sum / 10
    return
}

// global variable
var c, python, java bool = true, false, true

// constants
const (
    Pi = 3.14
    s = "hello world"
)

func main() {
    fmt.Println("Hello, Golang!")
    fmt.Println("rand number: ", rand.Intn(10))
    fmt.Printf("rand number: %g\n", math.Nextafter(2, 3))
    fmt.Println(add(2, 3))
    fmt.Println(swap("hello", "world"))
    fmt.Println(split(111))

    // local variable
    var i int = 1000

    // short declaration
    k := 3
    s := "hello"

    fmt.Println(i, c, python, java, k, s)

    const f = "%T(%v)\n"
    var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 -1
	z complex128 = cmplx.Sqrt(-5 + 2i)
    )
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
    
    fmt.Printf(f, Pi, Pi)
    fmt.Printf(f, s, s)
}
