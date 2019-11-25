import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
    res := make([][]string, len(searchWord))
    for i := range searchWord {
        res[i] = make([]string, 0)
    }
    for i := range searchWord {
        slist := make([]string, 0)
        for _, product := range products {
            if i+1 <= len(product) && searchWord[:i+1] == product[:i+1] {
                slist = append(slist, product)
            }
        }
        sort.Sort(sort.StringSlice(slist))
        if len(slist) > 3 {
            slist = slist[0:3]
        }
        res[i] = slist
    }
    return res
}
