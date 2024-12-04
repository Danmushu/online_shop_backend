package global

import (
	"strconv"
	"strings"
)

type Version string

func isNumeric(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func (v Version) ToInt() int {
	str := strings.ReplaceAll(string(v), ".", "")
	if len(str) != 3 || !isNumeric(str) {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

type VersionCode string

func (v VersionCode) ToInt() int {
	str := strings.ReplaceAll(string(v), "_", "")
	if len(str) != 8 || !isNumeric(str) {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
