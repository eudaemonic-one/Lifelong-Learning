func partition(s string) [][]string {
    path := make([]string, 0)
    ans := make([][]string, 0)
    backtrack(&ans, path, s)
    return ans
}

func backtrack(ans *[][]string, path []string, s string) {
    if (len(s) == 0) {
        *ans = append(*ans, path)
        return
    }
    for i := 0; i < len(s); i++ {
        if (isPalindrome(s[:i+1])) {
            path = append(path, s[:i+1])
            backtrack(ans, path, s[i+1:])
            tmp := make([]string, len(path)-1)
            copy(tmp, path)
            path = tmp
        }
    }
}

func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i < j {
        if (s[i] != s[j]) {
            return false
        }
        i += 1
        j -= 1
    }
    return true
}
