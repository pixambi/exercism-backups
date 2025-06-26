package strand


import "strings"

func ToRNA(dna string) string {
    var complement strings.Builder
    complements := map[rune]rune {
        'G': 'C',
        'T': 'A',
        'A': 'U',
        'C': 'G',
    }
    for _, v := range dna{
        complement.WriteRune(complements[v])
    }
    return complement.String()
}
