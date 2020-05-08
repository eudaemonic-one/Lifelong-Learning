func getHint(secret string, guess string) string {
    a, b := 0, 0
    secretDict := make(map[byte]int)
    guessDict := make(map[byte]int)
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            a++
        } else {
            secretDict[secret[i]]++
            guessDict[guess[i]]++
        }
    }
    for k, vs := range secretDict {
        if vg, ok := guessDict[k]; ok {
            b += min(vg, vs)
        }
    }
    return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
