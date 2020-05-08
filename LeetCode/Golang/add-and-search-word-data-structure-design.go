type WordDictionary struct {
    Children map[rune]*WordDictionary
    IsEnd bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{make(map[rune]*WordDictionary), false}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    cur := this
    for _, c := range word {
        if _, ok := cur.Children[c]; !ok {
            cur.Children[c] = &WordDictionary{make(map[rune]*WordDictionary), false}
        }
        cur = cur.Children[c]
    }
    cur.IsEnd = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    return backtrack(word, this)
}

func backtrack(word string, p *WordDictionary) bool {
    if len(word) == 0 {
        if !p.IsEnd {
            return false
        }
        return true
    }
    if word[0] == '.' {
        for _, child := range p.Children {
            if backtrack(word[1:], child) {
                return true
            }
        }
    } else if _, ok := p.Children[rune(word[0])]; ok {
        return backtrack(word[1:], p.Children[rune(word[0])])
    }
    return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
