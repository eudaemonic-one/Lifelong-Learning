func isHappy(n int) bool {
    slow := n
    fast := n
    for {
        slow = digitSquareSum(slow)
        fast = digitSquareSum(fast)
        fast = digitSquareSum(fast)
        if fast == 1 {
            return true
        }
        if slow == fast {
            break
        }
    }
    return false
}

func digitSquareSum(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}
