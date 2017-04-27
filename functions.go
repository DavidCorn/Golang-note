// multivalue return
func plusPlus(a, b, c int) (int, int)) {
    return a + b + c, a
}

// varadic functions
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
sum(1)
sum(2, 3, 4)

// closure
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}