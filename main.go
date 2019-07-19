package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Info struct {
		Xaid string `json:"xaid"`
	}
	
	
	jsonstring := `[{"xaid": "1000"}, {"xaid": "2000"}]`
	
	var test []Info
	json.Unmarshal([]byte(jsonstring), &test)
	fmt.Println(test)
}

