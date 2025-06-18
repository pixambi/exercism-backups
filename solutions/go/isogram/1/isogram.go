package isogram

import "unicode"


func IsIsogram(word string) bool {
    counts := make(map[rune]int)
    
	for _,v := range word {
        val := unicode.ToLower(v)
        _, exists := counts[val]
        if exists {
            return false
        }
        if(val != '-') && (val != ' '){
        	counts[val] += 1    
        }
    }
    return true
}
