package manager

import (
	"sync"
	"fmt"
)

type manager struct {
	Url string
}

var sharedManager *manager
var once sync.Once

func Shared() *manager {
	once.Do(func () {
		sharedManager = &manager{Url : "https://api.coinmarketcap.com/v1/ticker/?convert=EUR"}
	})

	return sharedManager
}
