package utils

// SliceString returns the string sliced a to b.
func SliceString(s string, a int, b int) string {
	return string([]rune(s)[a:b])
}

// SnipString returns the string snip a to b.
func SnipString(s string, a int, b int) string {
	runed := []rune(s)
	return string(runed[0:a]) + string(runed[b:])
}

// SnipStringCovered returns the string snipped covered strings.
// eg: "I want to[wanna] sleep." -> "I want to sleep."
func SnipStringCovered(s string, start string, end string) string {
	startIndex := -1
	startRune := []rune(start)[0]
	endRune := []rune(end)[0]

	runed := []rune(s)

	for i := 0; i < len(runed); i++ {
		if runed[i] == startRune && startIndex == -1 {
			startIndex = i
			break
		}
	}

	for i := len(runed) - 1; i >= 0; i-- {
		if runed[i] == endRune {
			s = SnipStringCovered(SnipString(s, startIndex, i+1), start, end)
			break
		}
	}

	return s
}

// LenString returns length of the string
func LenString(s string) int {
	return len([]rune(s))
}

// ReplaceStringFromIndex returns string replaced $index's char to $s in base
func ReplaceStringFromIndex(base string, s string, index int) string {
	runed := []rune(base)
	return string(runed[0:index]) + s + string(runed[index+1:])
}

// ObtainStringCovered returns covered strings.
// eg: "I want to[wanna] sleep." -> "wanna"
func ObtainStringCovered(s string, start string, end string) string {
	startIndex := -1
	startRune := []rune(start)[0]
	endRune := []rune(end)[0]

	runed := []rune(s)

	for i := 0; i < len(runed); i++ {
		if runed[i] == startRune && startIndex == -1 {
			startIndex = i
			break
		}
	}

	for i := len(runed) - 1; i >= 0; i-- {
		if runed[i] == endRune {
			s = SliceString(s, startIndex+1, i)
			break
		}
	}

	return s
}

// StringSliceRemove returns string slice after removed its element.
func StringSliceRemove(s []string, index int) []string {
	newSlice := make([]string, len(s)-1)

	skipped := false
	for i := 0; i < len(s)-1; i++ {
		if i == index {
			skipped = true
		}

		if skipped {
			newSlice[i] = s[i+1]
		} else {
			newSlice[i] = s[i]
		}
	}

	return newSlice
}

// StringSliceFind returns the index of the target element.
func StringSliceFind(s []string, target string) int {
	for i, str := range s {
		if str == target {
			return i
		}
	}

	return -1
}
