func canPermutePalindrome(s string) bool {
    dict := make(map[rune]int)
    for _, c := range s {
        if _, ok := dict[c]; ok {
            dict[c] += 1
        } else {
            dict[c] = 1
        }
    }
    odd := 0
    for _, v := range dict {
        if v % 2 == 1 {
            odd++
        }
    }
    return odd <= 1
}
