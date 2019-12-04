func groupStrings(strings []string) [][]string {
    diffsMap := make(map[int][]string)
    for _, str := range strings {
        if len(str) == 0 {
            diffsMap[0] = []string{""}
            continue
        }
        diffSum, base, diff := len(str), 0, 0
        for i := 1; i < len(str); i++ {
            base = int(math.Pow(26, float64(i-1)))
            diff = int((str[i]-str[i-1]+26)%26)
            diffSum += base * diff
        }
        if _, ok := diffsMap[diffSum]; ok {
            diffsMap[diffSum] = append(diffsMap[diffSum], str)
        } else {
            diffsMap[diffSum] = []string{str}
        }
    }
    i := 0
    res := make([][]string, len(diffsMap))
    for _, strs := range diffsMap {
        res[i] = strs
        i++
    }
    return res
}
