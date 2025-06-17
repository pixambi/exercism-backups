// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
    leap := false
    if year % 4 == 0{
        leap = true
        if year % 100 == 0{
            leap = false
            if year % 400 == 0 {
                leap = true
            }
        }
    }
    return leap
	}
