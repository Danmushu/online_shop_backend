package types

import (
	"github.com/fatih/structs"
	"github.com/spf13/cast"
	"sort"
	"strconv"
	"strings"
)

func ToString(v any) string {
	return strings.TrimSpace(cast.ToString(v))
}

func ToMap(model any) map[string]any {
	return structs.Map(model)
}

func ToUint(v any) uint {
	return cast.ToUint(v)
}

func ToFloat64(v any) float64 {
	return cast.ToFloat64(v)
}

func ToID(v any) uint {
	return ToUint(v)
}

func ToIDs(v []int) string {
	sort.Ints(v)
	s := ""
	for _, id := range v {
		s += strconv.Itoa(id) + ","
	}
	return s[0 : len(s)-1]
}

func ToIDsArray(s string) []int {
	var ids []int
	ss := strings.Split(s, ",")
	for _, s := range ss {
		id, _ := strconv.Atoi(s)
		ids = append(ids, id)
	}
	return ids
}

func ToInt(v any) int {
	return cast.ToInt(v)
}
func ToInt64(v any) int64 {
	return cast.ToInt64(v)
}

func ToBool(v any) bool {
	return cast.ToBool(v)
}

func IDsToStrings(li []uint) []string {
	var keys []string
	for _, v := range li {
		keys = append(keys, ToString(v))
	}
	return keys
}

func StringInArray(s string, array []string) bool {

	if array == nil {
		return false
	}

	for _, v := range array {
		if s == v {
			return true
		}
	}
	return false
}
