package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{
		Name:    "Denilson",
		Surname: "Son",
		Year:    1998,
	}

	t, err := json.Marshal(useall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	str := `{"username":"Denilson","surname":"ssSon","created":2023}`
	jsonRecord := []byte(str)
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", temp)
	}
}
