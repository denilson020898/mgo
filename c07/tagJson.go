package main

import (
	"encoding/json"
	"fmt"
)

type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created,omitempty"`
}

type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"`
}

func main() {
	noEmptyVar := NoEmpty{
		Name:    "Denilson",
		Surname: "",
		// Year: 1123,
	}

	t, err := json.Marshal(noEmptyVar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	passVar := Password{
		Name:    "Denilson",
		Surname: "",
		Year:    0,
		Pass:    "haleluah",
	}

	t, err = json.Marshal(passVar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	passStr := `{"username":"bruh", "surname":"", "creationyear":0}`

	jsonRecord := []byte(passStr)
	temp := Password{}
	err = json.Unmarshal(jsonRecord, &temp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", temp)
	}
}
