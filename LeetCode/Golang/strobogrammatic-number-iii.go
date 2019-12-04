func strobogrammaticInRange(low string, high string) int {
    dict := map[byte]string{'0':"0", '1':"1", '6':"9", '8':"8", '9':"6"}
    len1, len2 := len(low), len(high)
    lb, _ := strconv.Atoi(low)
    ub, _ := strconv.Atoi(high)
    cnt := 0
    for l := len1; l <= len2; l++ {
        backtrack(dict, "", lb, ub, l, &cnt)
    } 
    return cnt
}

func backtrack(dict map[byte]string, path string, low, high, length int, cnt *int) {
    if length == 0 {
        if len(path) > 1 && path[0] == '0' {
            return
        }
        num, _ := strconv.Atoi(path)
        if low <= num && num <= high {
            *cnt = *cnt + 1
        }
        return
    }
    for c := range dict {
        if length % 2 == 1 {
            if string(c) == dict[c] {
                backtrack(dict, string(c)+path, low, high, length-1, cnt)
            }
        } else {
            backtrack(dict, string(c)+path+dict[c], low, high, length-2, cnt)
        }
    }
}
