package main

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

type Item struct {
	Category string
	Quantity int32
}

type Student struct {
	Name string
	Age  int64
}

func main() {
	data, err := os.ReadFile("./andrei")
	check(err)

	var decoded Student
	err = bson.Unmarshal(data, &decoded)
	check(err)

	fmt.Printf("Unmarshalled Struct:\n%+v\n", decoded)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() {
	dat, err := os.ReadFile("./andrei")

	if err != nil {
		panic(err)
	}

	unmarshalAndPrint(dat)
}

func createElement() {
	doc, err := bson.Marshal(bson.D{{"category", "plate"}, {"quantity", 6}})

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./andrei", doc, 0644)
	if err != nil {
		panic(err)
	}

	var test Item
	err = bson.Unmarshal(doc, &test)
	fmt.Printf("Unmarshalled Struct:\n%+v\n", test)
}

func unmarshalAndPrint(data []byte) {
	var result Item
	err := bson.Unmarshal(data, &result)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshalled Struct:\n%+v\n", result)
}
