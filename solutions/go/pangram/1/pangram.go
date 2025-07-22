package pangram

import (
    "strings"
	"regexp"    
)

func IsPangram(input string) bool {
    letters := map[rune]bool{}
    reg := regexp.MustCompile("[^a-zA-Z]")
    cleaned := reg.ReplaceAllString(input, "")
    cleaned = strings.ToLower(cleaned)
    
    for _, letter := range cleaned {
        letters[letter] = true
    }
    
    return len(letters) == 26
}
