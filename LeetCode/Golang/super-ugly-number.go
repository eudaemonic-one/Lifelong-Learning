func nthSuperUglyNumber(n int, primes []int) int {
    const INT_MAX = int(^uint(0) >> 1)
    uglies := []int{1}
    cnt := make([]int, len(primes))
    for i := 1; i < n; i++ {
        ugly := INT_MAX
        for j := 0; j < len(primes); j++ {
            ugly = min(ugly, uglies[cnt[j]] * primes[j])
        }
        uglies = append(uglies, ugly)
        for j := 0; j < len(primes); j++ {
            if ugly == uglies[cnt[j]] * primes[j] {
                cnt[j]++
            }
        }
    }
    return uglies[n-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
