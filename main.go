package main

import (
	"booking_app/mapper"
	"booking_app/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	//var array = [100]string{}
	for {
		var arr []string
		var n int
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			sc.Scan()
			temps := sc.Text()
			arr = append(arr, temps)
		}
		updated := util.UpdateFunc(arr)
		mappedValues := mapper.MakeMap(updated)
		for index := range mappedValues {
			firstName, lastName, sl := mapper.GetMapValues(mappedValues, index)
			fmt.Printf("at %v : %v , %v created at %v\n", index, firstName, lastName, sl)
		}
	}
}
