package mapper

func GetMapValues(inpMap []map[string]string, index int) (string, string) {
	return inpMap[index]["firstName"], inpMap[index]["lastName"]
}
