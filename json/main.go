package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"],"Sample":[{"test":1},{"test 2":2}]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", f)

	m := f.(map[string]interface{})

	fmt.Printf("\n%v", m)

	s := m["Sample"].([]interface{})

	fmt.Printf("--\ns[0]%v", s[0])

}
