package main

import (
    "io"
    "fmt"
    "math"
    "os"
    "time"
    "strings"
)

type Vertex struct {
    x, y float64
}

type MyFloat float64

// Method on value type receiver
func (f MyFloat) Abs() float64 {
    if f < 0 {
	return float64(-f)
    }

    return float64(f)
}

// Method on pointer type receiver
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.x * v.x + v.y * v.y)
}


// Interface
type Abser interface {
    Abs() float64 
}

func (v *Vertex) Scale(s float64) {
    v.x = v.x * s
    v.y = v.y * s
}

func main() {
    v := &Vertex{3, 4} 
    fmt.Println(v.Abs())

    f := MyFloat(-3.14)
    fmt.Println(f.Abs())
    
    v.Scale(5)
    fmt.Println(v, v.Abs())

    var a Abser
    a = v
    fmt.Println(a.Abs())

    a = f
    fmt.Println(a.Abs())

    var rw ReadWriter
    rw = os.Stdout
    fmt.Fprintf(rw, "hello world\n")

    pw := Person{"WuXibin", 26}
    ps := Person{"TianWei", 22}

    fmt.Println(pw) 
    fmt.Println(ps) 

    addrs := map[string]IPaddr {
	"loopback" : {127, 0, 0, 1},
	"googledns" : {8, 8, 8, 8},
    }

    for n, a := range addrs {
	fmt.Printf("%v, %v\n", n, a)
    }

    if err := run(); err != nil {
	fmt.Println(err)
    }

    if v, e := Sqrt(5); e != nil {
	fmt.Println("ERROR", e)
    } else {
	fmt.Println("RESULT", v)
    }

    if v, e := Sqrt(-5); e != nil {
	fmt.Println("ERROR", e)
    } else {
	fmt.Println("RESULT", v)
    }

    r := strings.NewReader("Hello, reader!")
    b := make([]byte, 8)
    for {   // while
	n, e := r.Read(b)	
	fmt.Printf("n = %v, b = %v, e = %v\n", n, b, e)
	fmt.Printf("b[:n] == %q\n", b[:n])
	if e == io.EOF {
	    break
	}
    }

    reader := MyReader{}
    if n, e := reader.Read(b); e == nil {
	fmt.Printf("b[:n] == %q\n", b[:n])
    }

    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    rs := rot13Reader{s}
    io.Copy(os.Stdout, &rs)
}


// Interface
type Reader interface {
    Read(b []byte) (n int, err error)
}

type Writer interface {
    Write(b []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

type Person struct {
    Name string
    Age int
}

// String()
// fmt package look for this interface to print values [%v]
func (p Person) String() string {
    return fmt.Sprintf("%v (%v year)", p.Name, p.Age) 
}

type IPaddr [4]byte

func (ip IPaddr) String() string {
    return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}


// Error()
type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%v, %s", e.When, e.What)
}

func run() error {
    return &MyError{time.Now(), "it didn't work!"}	// return pointer
}


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("can't sqrt negative number: %f", float64(e))
}


func Sqrt(x float64) (float64, error) {
    if x < 0 {
	return 0, ErrNegativeSqrt(x)			// return value
    }

    return math.Sqrt(x), nil 
}


// Read()
type MyReader struct {}

func (r *MyReader) Read(b []byte) (int, error) {
    sz := len(b)
    for i := 0; i < sz; i++ {
	b[i] = 'A'
    }

    return sz, nil
} 

type rot13Reader struct {
    r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
    n, err = r.r.Read(b)

    for i := 0; i < n; i++ {
	if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
	    b[i] += 13
	} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
	    b[i] -= 13
	}    
    }

    return 
}
