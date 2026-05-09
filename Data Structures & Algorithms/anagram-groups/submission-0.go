import "slices"

func groupAnagrams(strs []string) [][]string {
    if len(strs) <= 1 {
        return [][]string{strs}
    }

    group := make([][]string, 0, 1000)
    groupMap := make(map[string][]string, 101)

    for _, str := range strs {
        strRunes := []rune(str)
        slices.Sort(strRunes)
        newSortedStr := string(strRunes)
        groupMap[newSortedStr] = append(groupMap[newSortedStr], str)
    }

    for _, groupStr := range groupMap {
        group = append(group, groupStr)
    }

    return group
}

