func canConstruct(ransomNote string, magazine string) bool {
    magazineMap := make(map[rune]int32, len(magazine)/2)    

    for _, r := range magazine {
        magazineMap[r] = magazineMap[r] + 1
    }
    
    for _, r := range ransomNote {
        cnt, ok := magazineMap[r]
        if !ok || cnt <= 0 {
            return false
        }

        magazineMap[r] = cnt - 1
    }

    return true
}