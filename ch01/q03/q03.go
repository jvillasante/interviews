package q03

func ReplaceSpaces(str []rune, length int) {
	spaceCount := 0
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	newLength := length + spaceCount*2
	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[newLength-1] = '0'
			str[newLength-2] = '2'
			str[newLength-3] = '%'
			newLength -= 3
		} else {
			str[newLength-1] = str[i]
			newLength -= 1
		}
	}
}
