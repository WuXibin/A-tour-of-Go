package main

import (
    "fmt"
    "math"
    "strings"
)

// struct
type Vertex struct {
    x, y int
}

func main() {
    // pointer
    i, j := 1024, 2048
    p := &i
    fmt.Printf("%T(%v)(%v)\n", p, p, *p)

    p = &j
    *p = 4096
    fmt.Printf("%T(%v)(%v)\n", p, p, *p)
    fmt.Println(j)

    // struct
    v := Vertex{1, 2}
    v.x = 10
    fmt.Printf("%T(%v)\n", v, v)

    pv := &v
    pv.x = 100
    pv.y = 200
    fmt.Printf("%T(%v)(%v)\n", pv, pv, pv)

    // array
    var arr [3]string
    arr[0] = "hello"
    arr[2] = "world"
    fmt.Println(arr[0], arr[1], arr[2])

    // slice
    s := []int{1, 7, 11, 24, 32, 48}
    fmt.Println("s ==", s)
    for i := 0; i < len(s); i++ {
	fmt.Printf("s[%d]=%d\n", i, s[i])
    }

    fmt.Println("s[1:3]", s[1:3])
    fmt.Println("s[:3]", s[:3])
    fmt.Println("s[4:]", s[4:])

    // slice with specified length and capacity
    a := make([]int, 5)	    // len = 5, cap = 5
    PrintSlice("a", a)
    b := make([]int, 0, 5)  // len = 0, cap = 5
    PrintSlice("b", b)
    c := b[:2]		    // len = 2, cap = 5
    PrintSlice("c", c)
    d := b[2:5]		    // len = 3, cap = 3
    PrintSlice("d", d)

    var z []int
    if z == nil {
	fmt.Println("z is nil")
    }
    z = append(z, 1)
    PrintSlice("z", z)
    z = append(z, 1, 2)
    PrintSlice("z", z)
    z = append(z, 10, 20, 30)
    PrintSlice("z", z)

    // slice range
    for i, val := range z {
	fmt.Printf("z[%d]=%d\n", i, val)
    }
    for _, val := range z {
	fmt.Printf("%d\n", val)
    }

    // multple array
    fmt.Println(Pic(10, 10))

    // map
    // 1) without intial value
    m := make(map[string]Vertex)
    m["Bell Labs"] = Vertex{40, -56}
    fmt.Println(m["Bell Labs"])
    fmt.Println(m, len(m))

    // 2) with intial value
    loc := map[string]Vertex{
	"Google" : {123, 456},
	"Microsoft" : {321, 654},
    }
    fmt.Println(loc, len(loc))

    // insert
    loc["Facebook"] = Vertex{123, 456}
    fmt.Println(loc, len(loc))

    // delete
    delete(loc, "Microsoft")
    delete(loc, "Yahoo!")

    // search
    val, ok := loc["Google"]
    fmt.Println("Google", val, "Exist?", ok)

    val, ok = loc["Linkedin"]
    fmt.Println("Linked", val, "Exist?", ok)

    fmt.Println(WordCount("Nice to see you, bye bye"))

    // iterator
    fmt.Println("iterator: ")
    for k, v := range loc {
	fmt.Println(k, v)
    }

    // function
    hypot := func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
    }

    fmt.Println(hypot(3, 4))

    // function closure
    // A closure is a function value that references variables from outside its body.
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
	fmt.Println(
	    pos(i),
	    neg(-2*i),
	)
    }

    f := Fibonacci()
    for i := 0; i < 10; i++ {
	fmt.Println(f())
    }
}

func PrintSlice(s string, x []int) {
    fmt.Printf("%s len=%d, cap=%d, %v\n", s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]int {
    pic := make([][]int, dx)
    for i := 0; i < dx; i++ {
	pic[i] = make([]int, dy)
	for j := 0; j < dy; j++ {
	    pic[i][j] = i * j
	}
    }
    return pic
}

func WordCount(s string) map[string]int {
    arr := strings.Fields(s)
    m := make(map[string]int)
    for _, val := range arr {
	m[val]++	
    }
    return m
}


func adder() func(int) int {
    // each return func boud to its own sum variable.
    sum := 0
    return func(x int) int {
	sum += x
	return sum
    }
}

func Fibonacci() func() int {
    a, b := 0, 1
    return func() int {
	a, b = b, a + b
	return a
    }
}
