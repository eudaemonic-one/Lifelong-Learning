func isUgly(num int) bool {
    if num <= 0 {
        return false
    }
    for _, divisor := range [3]int{2,3,5} {
        for num % divisor == 0 {
            num /= divisor
        }
    }
    return num == 1
}
