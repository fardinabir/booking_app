package mapper

func GetMapValues(inpMap []UserInfo, index int) (string, string, int) {
	return inpMap[index].firstName, inpMap[index].lastName, inpMap[index].serial
}
