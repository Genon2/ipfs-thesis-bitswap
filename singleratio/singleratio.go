package singleratio

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type SingleRatio struct {
	CID     string
	DHT     bool
	BITSWAP bool
}

var singleInstance *singleratio

func getInstance() *singleratio {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &singleratio{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
