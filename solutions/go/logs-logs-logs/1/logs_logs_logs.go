package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
        if v == '❗'{
            return "recommendation"
        }
         if v == '🔍'{
            return "search"
        }
         if v == '☀'{
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    newString := ""
	for _, v := range log{
        if v == oldRune {
            newString += string(newRune)
        } else {
            newString += string(v)
        }
    }
    return newString
    
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0
    for _, _ = range log{
        count++
    }
    return count <= limit
}
