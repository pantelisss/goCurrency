package manager

import (
	"coin_currency/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type manager struct {
	Url   string
	Ids   []string
	Coins []models.Coin
}

var Ids []string

var sharedManager *manager
var once sync.Once

func Shared() *manager {
	once.Do(func() {
		sharedManager = &manager{Url: "https://api.coinmarketcap.com/v1/ticker/?convert=EUR&limit=10"}
	})

	return sharedManager
}

func (m *manager) GetJson() {
	resp, err := http.Get(m.Url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &m.Coins)
	fmt.Println("get:\n", m.Coins[0])

	for _, v := range m.Coins {
		m.Ids = append(m.Ids, v.Id)
	}

	fmt.Println("All ids:\n", m.Ids)

	defer resp.Body.Close()
}

func (m *manager) GetJsonAsync(completion func()) {
	// channel := make(chan bool)
	go func() {
		m.GetJson()
		// channel <- true
		completion()
	}()

	// finshed := <-channel
	// fmt.Println("chanel finished", finshed)
	// completion()
}

func (m *manager) JsonData() (d []byte) {
	data, _ := json.Marshal(m.Coins)

	return data
}
