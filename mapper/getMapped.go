package mapper

import "time"

func GetMapValues(inpMap []userInfo, index int) (string, string, int) {
	time.Sleep(10 * time.Second)
	return inpMap[index].firstName, inpMap[index].lastName, inpMap[index].serial
}
