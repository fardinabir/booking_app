package main

import (
	"booking_app/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	//var array = [100]string{}
	var arr []string
	var n int
	fmt.Scan(&n)
	if n >= 4 {
		fmt.Printf("Please input lees\n")
	}
	for i := 0; i < n; i++ {
		sc.Scan()
		temps := sc.Text()
		arr = append(arr, temps)
	}
	updated := util.UpdateFunc(arr)
	for index, el := range updated {
		fmt.Printf("at ind: %v , %v\n", index, el)
	}
}
