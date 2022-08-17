package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type user struct {
	Firstname, Lastname string
}

func encode(users *[]user) []byte {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(&users); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func decode(buffer []byte) any {
	buf := bytes.NewBuffer(buffer)
	dec := gob.NewDecoder(buf)
	users := []user{}

	if err := dec.Decode(&users); err != nil {
		panic(err)
	}

	return users
}

func main() {
	users := []user{
		{Firstname: "Rosalinda", Lastname: "Fergusso"},
		{Firstname: "Bruce", Lastname: "Wayne"},
		{Firstname: "Yao", Lastname: "Ming"},
	}

	encoded := encode(&users)
	decoded := decode(encoded)

	fmt.Println(encoded, len(encoded), cap(encoded))

	// convert interface to struct
	value, ok := decoded.([]user)
	if ok {
		fmt.Println(value)
	}
}