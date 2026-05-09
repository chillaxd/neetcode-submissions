func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    counter := make(map[rune]int, len(s))

    for _, c := range s {
        counter[c] += 1
    }
    for _, c := range t {
        counter[c] -= 1
    }

    for _, count := range counter {
        if count != 0 {
            return false
        }
    }

    return true
}
