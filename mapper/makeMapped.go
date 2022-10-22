package mapper

import (
	"strings"
)

func MakeMap(inpSlice []string) []map[string]string {
	var ansMap = make([]map[string]string, 0)

	for _, el := range inpSlice {
		splittedNames := strings.Split(el, " ")
		tempMap := make(map[string]string)
		tempMap["firstName"] = splittedNames[0]
		tempMap["lastName"] = splittedNames[len(splittedNames)-1]
		ansMap = append(ansMap, tempMap)
	}
	return ansMap
}
