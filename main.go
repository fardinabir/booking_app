package main

import (
	"booking_app/mapper"
	"booking_app/util"
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	//var array = [100]string{}
	for x := 0; x < 5; x++ {
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
			fmt.Printf("At %v : %v , %v created at %v\n", index, firstName, lastName, sl)
		}
		if len(updated) != 0 {
			wg.Add(1)
			go clearenceSig(updated)
		}
		fmt.Println("Now give new Input")
	}
	wg.Wait()
	fmt.Println("Now finshed All///")
}

func clearenceSig(Sig []string) {
	time.Sleep(10 * time.Second)
	fmt.Println("###############\nCompleted\n###############\n", Sig)
	wg.Done()
}
