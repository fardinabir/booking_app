package util

func UpdateFunc(inpSlice []string) []string {
	var resSlice []string
	for i := 0; i < len(inpSlice); i++ {
		resSlice = append(resSlice, inpSlice[i]+"Updated")
	}
	return resSlice
}
