func removeDuplicateLetters(s string) string {
    if len(s) == 0 {
        return ""
    }
    seen := make(map[byte]bool)
    lastOccur := make(map[byte]int)
    stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        lastOccur[s[i]] = i
    }
    for i := 0; i < len(s); i++ {
        ch := s[i]
        if _, ok := seen[ch]; !ok {
            for len(stack) > 0 && ch < stack[len(stack)-1] && i < lastOccur[stack[len(stack)-1]] {
                delete(seen, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            seen[ch] = true
            stack = append(stack, ch)
        }
    }
    return string(stack)
}

// Approach 1: Greedy - Solving Letter by Letter

// func removeDuplicateLetters(s string) string {
//     if len(s) == 0 {
//         return ""
//     }
//     counter := make([]int, 26)
//     for _, ch := range s {
//         counter[ch-'a']++
//     }
//     // In each iteration, we determine leftmost letter in our solution.
//     // This will be the smallest character such that
//     // its suffix contains at least one copy of every character in the string.
//     pos := 0 
//     for i := 0; i < len(s); i++ {
//         if s[i] < s[pos] {
//             pos = i
//         }
//         counter[s[i]-'a']--
//         // end the iteration once the suffix doesn't have each unique character
//         if counter[s[i]-'a'] == 0 {
//             break
//         }
//     }
//     substr := make([]byte, 0)
//     for i := pos; i < len(s); i++ {
//         if s[i] != s[pos] {
//             substr = append(substr, s[i])
//         }
//     }
//     return string(s[pos]) + removeDuplicateLetters(string(substr))
// }
