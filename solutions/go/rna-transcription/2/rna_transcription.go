package strand

func ToRNA(dna string) string {
	complement := ""
    complements := map[rune]rune {
        'G': 'C',
        'T': 'A',
        'A': 'U',
        'C': 'G',
    }
    for _, v := range dna{
        complement += string(complements[v])
    }
    return complement
}
