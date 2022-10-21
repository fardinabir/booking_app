package util

import "fmt"

func UpdateFunc(inpSlice []string) []string {
	fmt.Print("Module was called")
	resSlice := []string{}
	for i := 0; i < len(inpSlice); i++ {
		resSlice = append(resSlice, inpSlice[i]+"Updated")
	}
	return resSlice
}
