package repeat

func hasRepeatedA(word string) bool {
    for i := 0; i < len(word)-1; i++ {
        if word[i] == 'a' && word[i+1] == 'a' {
            return true
        }
    }
    return false
}