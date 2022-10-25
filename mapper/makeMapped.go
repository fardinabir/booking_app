package mapper

import (
	"strings"
)

type userInfo struct {
	firstName string
	lastName  string
	serial    int
}

func MakeMap(inpSlice []string) []userInfo {

	var ansMap = make([]userInfo, len(inpSlice))

	for ind, el := range inpSlice {
		splittedNames := strings.Split(el, " ")
		ansMap[ind].firstName = splittedNames[0]
		ansMap[ind].lastName = splittedNames[len(splittedNames)-2]
		ansMap[ind].serial = ind
	}
	return ansMap
}
