// variable declaration
// outside functions
var i int = 0
var s string // ""

var i, j, k
var b, f, s = true, 2.3, "four"

var f, err = os.open(name)

// inside function
// short variable declaration
t := 0.0
var boiling float64 = 100

// swap
i, j = j, i

// Pointers
p := &x // p, of type *int, points to x, aka the address of x
*p = 2

// built-in function new
p := new(int)

// type declarations
type Celsius float64 // package level, if starts with an upper-acse, accessible from other packages as well

// import
import (
	"fmt"
	"gop1.io/ch2/tempconv"
)

// package initiation
func init() {/* ... */} // cannot be referenced, it automatically executed
