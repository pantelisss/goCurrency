package manager

import (
	"sync"
	"net/http"
	"fmt"
	"io/ioutil"
	"coin_currency/models"
	"encoding/json"
)

type manager struct {
	Url string
}

var sharedManager *manager
var once sync.Once

func Shared() *manager {
	once.Do(func () {
		sharedManager = &manager{Url : "https://api.coinmarketcap.com/v1/ticker/?convert=EUR&limit=10"}
	})

	return sharedManager
}

func (m *manager) GetJson() {
	resp, err := http.Get(m.Url)
	if err != nil {
		fmt.Println(err)	
	}

	body, err := ioutil.ReadAll(resp.Body)
	var coins []models.Coin	
	json.Unmarshal(body,&coins)
	fmt.Println("get:\n", coins[0])
  	
  	var ids []string
  	for _, v := range coins {
  		ids = append(ids, v.Id) 		
  	} 	
  	
  	fmt.Println("All ids:\n", ids)

  	defer resp.Body.Close()
  }
