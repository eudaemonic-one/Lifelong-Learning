func canWin(s string) bool {
    consecutives := make([]int, 0)
    length := 0
    maxLen := 0
    for i := 0; i < len(s)-1; {
        length = 1
        for i < len(s)-1 && s[i] == '+' && s[i+1] == '+' {
            length++
            i++
        }
        if length > 1 {
            consecutives = append(consecutives, length)
            if length > maxLen {
                maxLen = length
            }
        }
        i++
    }
    g := make([]int, len(s)+1)
    for i := 2; i <= maxLen; i++ {
        gsub := make(map[int]int)
        for first := 0; first < i/2; first++ {
            second := i - first - 2
            gsub[g[first] ^ g[second]] += 1
        }
        g[i] = firstMissingNumber(gsub)
    }
    xorValue := 0
    for i := 0; i < len(consecutives); i++ {
        xorValue ^= g[consecutives[i]]
    }
    return xorValue != 0
}

func firstMissingNumber(set map[int]int) int {
    for i := 0; i < len(set); i++ {
        if _, ok := set[i]; !ok {
            return i
        }
    }
    return len(set)
}
