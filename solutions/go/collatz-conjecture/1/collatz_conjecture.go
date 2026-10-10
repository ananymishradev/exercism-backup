package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
	count := 0
    if n <= 0 {
        return count, errors.New("Error")
    }
    if n == 1 {
        return count, nil
    }
    for n != 1 {
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = (n * 3) + 1
        } 
        count ++
    }
    return count, nil
}
