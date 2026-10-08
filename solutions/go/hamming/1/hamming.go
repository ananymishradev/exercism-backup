package hamming
import "errors"
func Distance(a, b string) (int, error) {
    if a == b {
        return 0, nil
    }
    if len(a) != len (b) {
        return 0, errors.New("Unequal Length")
    }
    count := 0
    for i := 0 ; i < len(a) ; i++ {
        if a[i] != b[i] {
            count ++
        }
    }
    return count, nil
}