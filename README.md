```go
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

	/*output
	[13 255 131 2 1 2 255 132 0 1 255 130 0 0 45 255 129 3 1 1 4 117 115 101 114 1 255 130 0 1 2 1 9 70 105 114 115 116 110 97 109 101 1 12 0 1 8 76 97 115 116 110 97 109 101 1 12 0 0 0 53 255 132 0 3 1 9 82 111 115 97 108 105 110 100 97 1 8 70 101 114 103 117 115 115 111 0 1 5 66 114 117 99 101 1 5 87 97 121 110 101 0 1 3 89 97 111 1 4 77 105 110 103 0] 114 128
	
	[{Rosalinda Fergusso} {Bruce Wayne} {Yao Ming}]
	*/
}
```
