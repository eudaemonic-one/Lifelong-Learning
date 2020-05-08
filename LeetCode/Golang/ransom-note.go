func canConstruct(ransomNote string, magazine string) bool {
    dict := make(map[byte]int)
    for i := 0; i < len(magazine); i++ {
        dict[magazine[i]]++
    }
    for i := 0; i < len(ransomNote); i++ {
        dict[ransomNote[i]]--
    }
    for ch := range dict {
        if dict[ch] < 0 {
            return false
        }
    }
    return true
}
