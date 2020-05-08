func generatePossibleNextMoves(s string) []string {
    res := make([]string, 0)
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '+' && s[i+1] == '+' {
            flip := s[:i] + "--" + s[i+2:]
            res = append(res, flip)
        }
    } 
    return res
}
