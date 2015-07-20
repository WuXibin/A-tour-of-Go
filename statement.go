package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

func Sqrt(x float64) float64 {
    z := 1.0
    delta := 0.001

    newz := z - (z * z - x) / (2 * z)

    for i := 0; math.Abs(z - newz) > delta; i++ {
	z = newz
	newz = z - (z * z - x) / (2 * z)
	fmt.Println("loop:", i, z)
    }

    return z
} 

func main() {
    defer fmt.Println("hello")

    sum := 0

    // for loop
    for i := 0; i < 10; i++ {
	sum += i
    }

    for sum < 1000 {
	sum <<= 1
    }

    // if
    if sum += 2000; sum < 10000 {
	fmt.Println("sum less than 10000")
    }

    fmt.Println(sum)

    fmt.Println(Sqrt(1), math.Sqrt(1))
    fmt.Println(Sqrt(2), math.Sqrt(2))
    fmt.Println(Sqrt(3), math.Sqrt(3))
    fmt.Println(Sqrt(4), math.Sqrt(4))

    // switch
    switch os := runtime.GOOS; os {
	case "darwin":
	    fmt.Println("OS X")
	case "linux":
	    fmt.Println("Linux")
	default:
	    fmt.Printf("%s\n", os)
    }

    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
	fmt.Println("Today.")
    case today + 1:
	fmt.Println("Tomorrow.")
    case today + 2:
	fmt.Println("In two days.")
    default:
	fmt.Println("Too far away.")
    }

    fmt.Println("counting")

    for i := 0; i < 10; i++ {
	defer fmt.Println(i)
    }

    fmt.Println("done")

}
