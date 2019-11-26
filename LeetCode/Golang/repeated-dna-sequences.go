func findRepeatedDnaSequences(s string) []string {
    res := make([]string, 0)
    dict := make(map[string]int)
    for i := 0; i <= len(s)-10; i++ {
        if _, ok := dict[s[i:i+10]]; ok {
            dict[s[i:i+10]] += 1
        } else {
            dict[s[i:i+10]] = 1
        }
    }
    for seq := range dict {
        if dict[seq] > 1 {
            res = append(res, seq)
        }
    }
    return res
}
