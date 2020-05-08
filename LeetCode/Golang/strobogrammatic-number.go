func isStrobogrammatic(num string) bool {
    dict := map[byte]string{'0':"0", '1':"1", '2':"*", '3':"*", '4':"*", '5':"*", '6':"9", '7':"*", '8':"8", '9':"6"}
    rotated := ""
    for i := len(num)-1; i >= 0; i-- {
        if v, ok := dict[num[i]]; ok && v == "*" {
            return false
        }
        rotated += dict[num[i]]
    }
    return rotated == num
}
