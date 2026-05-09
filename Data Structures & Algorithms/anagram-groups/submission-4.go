func groupAnagrams(strs []string) [][]string {
    if len(strs) <= 1 {
        return [][]string{strs}
    }

    group := make([][]string, 0, 1000)
    groupMap := make(map[[26]int][]string, 101)

    for _, str := range strs {
        var counts [26]int
        for i:=0;i<len(str);i++ {
            counts[str[i]-'a']++
        }

        groupMap[counts] = append(groupMap[counts], str)
    }

    for _, groupStr := range groupMap {
        group = append(group, groupStr)
    }

    return group
}
