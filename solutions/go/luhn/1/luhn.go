package luhn

import "strings"


func Valid(id string) bool {
    trimmed := strings.Fields(strings.TrimSpace(id))
	clean := strings.Join(trimmed, "")
    for _, char := range clean {
    if char < '0' || char > '9' {
        return false
    }
}
    length := len(clean)
	if length <= 1{
        return false
    }
	
    sum := 0
    for i := length -1; i >= 0; i--{
    val := int(clean[i] - '0')
    if (length - 1 - i) % 2 == 1 { 
        val = val * 2
        if val > 9 {
            val = val - 9
        }
    }
    sum += val
}
    if sum % 10 == 0 {
        return true
    }
    return false
}
