package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	arrays := [][]int{
		{1, 1, 1, 1},
		{4, 2, 3, 4},
		{4, 2, 3, 4},
		{4, 2, 3, 4},
	}

	for _, r := range arrays {
		fmt.Println(r)
	}
	fmt.Println()
	fmt.Println(GetTheMostPopular(arrays))
}
func GetTheMostPopular(arrays [][]int) (final []int) {

	arrayMap := make(map[string]int)
	for _, array := range arrays {

		b, _ := json.Marshal(array)
		s := string(b)
		arrayMap[s]++

	}
	s := ""
	for k, v := range arrayMap {
		if s == "" {
			s = k
		}

		if arrayMap[s] < v {
			s = k
		}
	}
	json.Unmarshal([]byte(s), &final)
	return
}
