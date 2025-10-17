func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sRunes := []rune(s)
    tRunes := []rune(t)
    
    sRuneMap := make(map[rune]int32, len(sRunes))
    tRuneMap := make(map[rune]int32, len(tRunes))

    for _, r := range sRunes {
       sRuneMap[r] = sRuneMap[r] + 1
    }

    for _, r := range tRunes {
       tRuneMap[r] = tRuneMap[r] + 1
    }
    
    for r, cnt := range sRuneMap {
        tCnt, ok := tRuneMap[r] 
        if !ok {
            return false
        }
        
        if cnt != tCnt {
            return false
        }
    }

    return true
}