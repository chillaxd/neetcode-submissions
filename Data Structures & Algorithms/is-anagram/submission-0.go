import "slices"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sChars, tChars := []rune(s), []rune(t)

    slices.Sort(sChars)
    slices.Sort(tChars)

    return string(sChars) == string(tChars)
}
