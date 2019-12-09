func wordPattern(pattern string, str string) bool {
    words := strings.Split(str, " ")
    if len(words) != len(pattern) {
        return false
    }
    dict1 := make(map[byte]string)
    dict2 := make(map[string]byte)
    for i := 0; i < len(words); i++ {
        word := words[i]
        pat := pattern[i]
        if val, ok := dict1[pat]; ok && val != word {
            return false
        } else {
            dict1[pat] = word
        }
        if val, ok := dict2[word]; ok && val != pat {
            return false
        } else {
            dict2[word] = pat
        }
    }
    return true
}
