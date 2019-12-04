func findStrobogrammatic(n int) []string {
    dict := map[byte]string{'0':"0", '1':"1", '6':"9", '8':"8", '9':"6"}
    output := make([]string, 0)
    backtrack(dict, "", n, &output)
    return output
}

func backtrack(dict map[byte]string, path string, n int, output *[]string) {
    if n == 0 {
        if len(path) > 1 && path[0] == '0' {
            return
        }
        *output = append(*output, path)
        return
    }
    for c := range dict {
        if n % 2 == 1 {
            if string(c) == dict[c] {
                backtrack(dict, string(c)+path, n-1, output)
            }
        } else {
            backtrack(dict, string(c)+path+dict[c], n-2, output)
        }
    }
}
