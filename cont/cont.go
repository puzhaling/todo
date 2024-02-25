package cont

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
}

func SaveItems(filename string, 
items []Item) error {

	jsonArr, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(filename, jsonArr, 0664)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	x, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, nil
	}
	
	var items []Item
	if err := json.Unmarshal(x, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}