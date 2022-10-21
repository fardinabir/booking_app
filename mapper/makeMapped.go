package mapper

import "strings"

func MakeMap(inpSlice []string) []map[string]string {
	var ansMap = make([]map[string]string, 0)

	for ind, el := range inpSlice {
		splittedNames := strings.Split(el, " ")
		ansMap[ind]["firstName"] = splittedNames[0]
		ansMap[ind]["lastName"] = splittedNames[len(splittedNames)-1]
	}
	return ansMap
}
