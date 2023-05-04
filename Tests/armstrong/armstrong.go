package armstrong

import (
	"fmt"
)

func isArmstrongNumber(n int) bool {
    numDigits := len(fmt.Sprint(n))
    sum := 0
    for i := n; i > 0; i /= 10 {
        digit := i % 10
        sum += pow(digit, numDigits)
    }
    return sum == n
}

func pow(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * pow(x, y-1)
}