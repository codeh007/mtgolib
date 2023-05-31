package utils

import (
	"regexp"
	"strings"
)

/*
字符串数组去重
时间复杂度：O(n)
空间复杂度：O(1)
*/
func RemoveDuplication_sort[T any](arr []*T) []*T {
	length := len(arr)
	if length == 0 {
		return arr
	}

	j := 0
	for i := 1; i < length; i++ {
		if arr[i] != arr[j] {
			j++
			if j < i {
				swap(arr, i, j)
			}
		}
	}

	return arr[:j+1]
}

func swap[T any](arr []*T, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

// wildCardToRegexp converts a wildcard pattern to a regular expression pattern.
func wildCardToRegexp(pattern string) string {
	components := strings.Split(pattern, "*")
	if len(components) == 1 {
		// if len is 1, there are no *'s, return exact match pattern
		return "^" + pattern + "$"
	}
	var result strings.Builder
	for i, literal := range components {

		// Replace * with .*
		if i > 0 {
			result.WriteString(".*")
		}

		// Quote any regular expression meta characters in the
		// literal text.
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return "^" + result.String() + "$"
}

/*
whildcard 字符匹配
例子：

	WildcardMatch("t*e*s*t", "ttttteeeeeeeesttttttt") //true
	WildcardMatch("t*e*s*t", "tset") //false
*/
func WildcardMatch(pattern string, value string) bool {
	result, _ := regexp.MatchString(wildCardToRegexp(pattern), value)
	return result
}

/*
`abc;ddd` -> `["abc","ddd"]`
*/
func CustomSplitStr(input string) []string {
	a := strings.FieldsFunc(input, customDelimitersSplit)
	return removeEmptyStrings(a)
}

func customDelimitersSplit(r rune) bool {
	return r == ';' || r == '|' || r == ' ' || r == '\n' || r == '\t'
}

// removeEmptyStrings - Use this to remove empty string values inside an array.
// This happens when allocation is bigger and empty
func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
