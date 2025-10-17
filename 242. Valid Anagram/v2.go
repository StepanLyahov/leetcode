func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sRunes := []rune(s)
    tRunes := []rune(t)
    
    sRuneMap := make(map[rune]int32, len(sRunes))

    for _, r := range sRunes {
       sRuneMap[r] = sRuneMap[r] + 1
    }

    for _, r := range tRunes {
        cnt, ok := sRuneMap[r]
        if !ok {
            return false
        }

        if cnt <= 0 {
            return false
        }

        sRuneMap[r] = cnt - 1
    }
    
    for _, cnt := range sRuneMap {
        if cnt != 0 {
            return false
        }
    }

    return true
}