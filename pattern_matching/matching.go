package pattern

import "errors"

func skipChars(startPos int, text *[]rune, runCheck rune, previous *rune) (int, error) {
	k := 0
	for k = startPos; k < len(*text); k++ {
		current := (*text)[k]
		if current == runCheck {
			return k - 1, nil
		}
		if previous != nil && current != *previous {
			return 0, errors.New("pattern mismatch")
		}
	}
	return k, nil
}

func IsMatch(text string, pattern string) bool {
	patternArray := []rune(pattern)
	textArray := []rune(text)
	if len(pattern) == 0 || len(text) == 0 {
		return false
	}
	if patternArray[0] == '*' {
		return false
	}
	validPattern := false
	patternIndex := 0
	nextPattern := rune(0)
	previousPattern := rune(0)
	currentPattern := rune(0)
	for valueIndex := 0; valueIndex < len(textArray); valueIndex++ {
		current := textArray[valueIndex]
		if patternIndex < len(patternArray) {
			previousPattern = currentPattern
			currentPattern = patternArray[patternIndex]
			if patternIndex < len(patternArray)-1 {
				nextPattern = patternArray[patternIndex+1]
			}
			if current == currentPattern {
				validPattern = true
				patternIndex++
			} else if currentPattern == '*' {
				// aaabb*
				if patternIndex == len(patternArray)-1 {
					return true
				}
				idx, mismatchError := skipChars(valueIndex, &textArray, nextPattern, &previousPattern)
				if mismatchError != nil {
					return false
				}
				valueIndex = idx
				patternIndex++
			} else if currentPattern == '.' {
				if nextPattern == rune(0) {
					return true
				}
				valueIndex, _ = skipChars(valueIndex, &textArray, nextPattern, nil)
				patternIndex++
			} else {
				validPattern = false
				break
			}
		} else {
			validPattern = false
		}
	}
	return validPattern
}
