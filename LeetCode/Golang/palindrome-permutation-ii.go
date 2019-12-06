func generatePalindromes(s string) []string {
    perms := make([]string, 0)
    flag, dict := canPermutePalindrome(s)
    if !flag {
        return perms
    }
    backtrack(dict, "", &perms)
    return perms
}

func canPermutePalindrome(s string) (bool, map[rune]int) {
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
    return odd <= 1, dict
}

func backtrack(dict map[rune]int, pal string, perms *[]string) {
    // check if there is remaining character in dict
    remain := 0
    for _, n := range dict {
        remain += n
    }
    if remain == 0 && len(pal) > 0 {
        // check if pal is a palindrome
        i, j := 0, len(pal)-1
        for i < j {
            if pal[i] != pal[j] {
                return
            }
            i++
            j--
        }
        *perms = append(*perms, pal)
        return
    }
    for c, n := range dict {
        if n <= 0 {
            continue
        }
        if n % 2 == 0 {
            dict[c] -= 2
            backtrack(dict, string(c)+pal+string(c), perms)
            dict[c] += 2
        } else {
            dict[c] -= 1
            backtrack(dict, string(c)+pal, perms)
            dict[c] += 1
        }
    }
}
