// integers
// int8, 16, 32, 64
// uint8, 16, 32, 64

// float-point numbers
// float32, float64

//const
const e = 2.7

// complex numbers
// complex64, complex128
x := 1 + 2i
var y complex64 = complex(1, 2)

// boolean
var flag bool = true

// strings
s := "hello world"
//iterate through s
func iterateS(name string) string{
	for i := 0; i < len(name); i++ {
		//...
	}
	return name
}

// convert byte to string
s := "abc"
b := []byte(s)
s2 := string(b)

// convert int to string
s = strconv.Itoa(123) // "123"
// convert string to int
x, err = strconv.Atoi("123") // 123
y, err = strconv.PaseInt("123", 10, 64)

// contant generator
const (
	one int = Itoa
	two
	three
)

// arrays
var a [3]int
// for loop
for i, v := range a {
	fmt.Printf("%d, %d\n", i, v)
}

	q := [...]int{1, 2, 3}
	q1 := [...]int{99: -1
}

// mutable variable pass
func zero(ptr *[32]byte) {/*...*/}

// slices
s := make([]string, 3)
s[0] = "a"
s[1] = "b"
s[2] = "c"
s = append(s, "d")

len(s) // actual length of the slice
cap(s) // start from the slice, end at the actual length of the base array

// test whether a slice is empty, use len, don't use cap!
var s []int
s = nil // len(s) = 0, cap(s) = nil
s = []int{nil} // len(s) = 0, cap(s) = nil
s = []int{} // len(s) = 0, cap(s) != nil

// append function
/*
each time append function check whether the capacity is sufficient,
otherwise double the capacity (e.g. make([]int, len(s), 2 * cap(s); copy(z, x))
*/
var runes []string
runes = append(runes, 'r')

// implementing stack
stack := make([]int, 5)
stack = append(stack, v)
// top of stack
top := stack[len(stack) - 1]
// pop
stack = stack[: len(stack) - 1]

// Map, Set, struct: http://www.jianshu.com/p/247ba63ad8db


