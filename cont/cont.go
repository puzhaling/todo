package cont

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text string
	Priority int
	position int
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

	for i, _ := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func (item *Item) SetPriority(priority int) {
	switch priority {
	case 1:
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func (item *Item) PrettyP() string {
	switch item.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"
	default:
		return " "
	}
}

func (item *Item) Label() string {
	return strconv.Itoa(item.position) + "."
}


/* 
	ByPri implements sort.Interface for []Item based on
	the Priority & position fields
*/
type ByPri []Item

func (s ByPri) Len() int { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}