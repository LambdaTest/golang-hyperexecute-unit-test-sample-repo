package integer

func lenInt(n int) int {
    // function to calculate the length of an integer
    if n == 0 {
        return 1
    }

    count := 0
    for n != 0 {
        n /= 10
        count++
    }

    return count
}

func addInt(a, b int) int {
    // function to add two integers
    return a + b
}