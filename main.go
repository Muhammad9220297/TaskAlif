package main

import (
	"AlifTack/Senders"
	"AlifTack/database"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if err := database.Connect(); err != nil {
		fmt.Println(" No Connect base!", err)
	} else {
		fmt.Println("Connect")
		p := Senders.Pakupkas{}
		if err = p.ScanTable(); err != nil {
			fmt.Println("No Scan ", err)
		} else {
			//wg.Add(1)
			Senders.Selector(p)
			//	wg.Wait()

		}

	}

}
