// Package jpnumber provides a function to convert integers to their Japanese number representations.
package jpnumber

import (
	"fmt"
	"strings"
)

var (
	kanjiDigits = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	kanjiUnits  = []string{"", "十", "百", "千", "万", "億", "兆", "京"}
)

// ToJPNumber converts an integer to its Japanese number representation.
// It returns the Japanese number as a string.
// The maximum number that can be converted is 99999999999999999.
func ToJPNumber(num int) (string, error) {
	if num == 0 {
		return kanjiDigits[0], nil
	}

	if num > 99_999_999_999_999_999 {
		return "", fmt.Errorf("number is too large: %d. max is 99999999999999999", num)
	}

	var jpNumber []string
	unitIndex := 0
	for num > 0 {
		jpNumber, _ = appendDigitKanji(num, unitIndex, jpNumber)
		num /= 10
		unitIndex++

		if unitIndex >= 4 && num > 0 {
			for j := 0; num > 0; j++ {
				jpNumber = append([]string{kanjiUnits[unitIndex]}, jpNumber...)
				isAppended := false

				// 4桁ごとに処理
				for subUnitIndex := range []int{0, 1, 2, 3} {
					if num <= 0 {
						break
					}

					jpNumber, isAppended = appendDigitKanji(num, subUnitIndex, jpNumber)
					num /= 10
				}

				if jpNumber[len(jpNumber)-1] == kanjiUnits[unitIndex] && !isAppended {
					jpNumber = jpNumber[:len(jpNumber)-1]
				}

				unitIndex++
			}
		}
	}

	return strings.Join(jpNumber, ""), nil
}

func appendDigitKanji(num int, unitIndex int, jpNumber []string) ([]string, bool) {
	isAppended := false
	digit := num % 10
	if digit > 0 {
		digitKanji := kanjiDigits[digit]
		if unitIndex != 0 && digit == 1 {
			digitKanji = ""
		}

		jpNumber = append([]string{digitKanji + kanjiUnits[unitIndex]}, jpNumber...)
		isAppended = true
	}

	return jpNumber, isAppended
}
