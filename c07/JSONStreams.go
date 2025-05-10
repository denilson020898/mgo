package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}

	return temp
}

func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}
	return err
}

func JSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func main() {
	var i int
	var t Data
	for i = 0; i < 5; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	fmt.Println("Last record:", t)
	_ = PrettyPrint(t)

	val, _ := JSONstream(DataRecords)
	fmt.Println(val)

	buf := new(bytes.Buffer)

	encoder := json.NewEncoder(buf)
	err := Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("After Serialize:", buf)

	decoder := json.NewDecoder(buf)
	var temp []Data
	err = DeSerialize(decoder, &temp)
	fmt.Println("After DeSerialized:")
	for index, value := range temp {
		fmt.Println(index, value)
	}
}
