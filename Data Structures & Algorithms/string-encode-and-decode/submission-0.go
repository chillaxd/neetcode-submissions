
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var output string
	for _, str := range strs {
		strLen := len(str)
		if strLen == 0 {
			output += "1000"
			continue
		}

		chars := []rune(str)
		n := rand.Intn(strLen) + 1
		m := rune(n)

		for i, c := range chars {
			chars[i] = c + m
		}

		encodedStr := fmt.Sprintf("%d%s",1200-n,string(chars)) 

		output += fmt.Sprintf("%d%s",1000+len(encodedStr),encodedStr) 
	}

	return output
}

func (s *Solution) Decode(encoded string) []string {
	output := []string{}

	for {
		if len(encoded) == 0 {
			break
		}

		ls := encoded[:4]
		if ls == "1000" {
			output = append(output, "")
			encoded = encoded[4:]
			continue
		}

		l, err := strconv.Atoi(ls)
		if err != nil {
			return []string{}
		}
		strLen := l - 1000
		decodedStrWithN := encoded[4:4+strLen]

		modN := decodedStrWithN[:4]
		modn, err := strconv.Atoi(modN)
		if err != nil {
			return []string{}
		}
		n := 1200 - modn
		m := rune(n) 
		chars := []rune(decodedStrWithN[4:])

		for i, c := range chars {
			chars[i] = c - m
		}
		
		output = append(output, string(chars))
		encoded = encoded[4+strLen:]
	}

	return output
}
