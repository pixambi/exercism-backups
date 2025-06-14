package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return -1, errors.New("must be equal length")
    }
    count := 0
    for i := 0; i < len(a); i++{
        if a[i] != b[i]{
            count++
        }
    }
   
    return count, nil
}
