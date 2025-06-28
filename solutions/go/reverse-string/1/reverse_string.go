package reverse

import "strings"

func Reverse(input string) string {
    var builder strings.Builder
    runes := []rune(input)
	for i := len(runes) -1; i >= 0 ; i--{
        builder.WriteRune(runes[i])
    }
    return builder.String()
}
