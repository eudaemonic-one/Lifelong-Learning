func isPalindrome(s string) bool {
    var l, r = 0, len(s)-1
    for l < r {
        for l < r && !isAlphaNum(s[l]) {
            l += 1
        }
        for l < r && !isAlphaNum(s[r]) {
            r -= 1
        }
        if !(s[l] == s[r] || (isAlpha(s[l]) && isAlpha(s[r]) && (s[l] - s[r] == 32 || s[r] - s[l] == 32))) {
            return false
        }
        l += 1
        r -= 1
    }
    return true
}

func isAlphaNum(c byte) bool {
    if (c < '0' || (c > '9' && c < 'A') || (c > 'Z' && c < 'a') || c > 'z') {
        return false
    }
    return true
}

func isAlpha(c byte) bool {
    if (c < 'A' || (c > 'Z' && c < 'a') || c > 'z') {
        return false
    }
    return true
}