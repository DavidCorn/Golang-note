// for loop
// A classic initial/condition/after `for` loop.
for j := 7; j <= 9; j++ {
    fmt.Println(j)
}

// while loop
for {
    fmt.Println("loop")
    break
}

// if condition
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}

// switch
// Here's a basic `switch`.
i := 2
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}

// You can use commas to separate multiple expressions
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}

// to express if/else logic. Here we also show how the
// `case` expressions can be non-constants.
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon")
default:
    fmt.Println("It's after noon")
}

// A type `switch` compares types instead of values.  You
// can use this to discover the the type of an interface
// value.  In this example, the variable `t` will have the
// type corresponding to its clause.
whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int")
    default:
        fmt.Printf("Don't know type %T\n", t)
    }
}