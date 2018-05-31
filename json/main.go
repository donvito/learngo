package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"],"Sample":[{"test":1},{"test 2":2}]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)

	fmt.Printf("%v", f)

	if err != nil {
		panic(err)

	}
	m := f.(map[string]interface{})

	fmt.Printf("\n%v", m)

	s := m["Sample"].([]interface{})

	fmt.Printf("--\ns[0]%v", s[0])

	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Println(k, "is string", vv)
	// 	case float64:
	// 		fmt.Println(k, "is float64", vv)
	// 	case []interface{}:
	// 		fmt.Println(k, "is an array:")
	// 		for i, u := range vv {
	// 			fmt.Println(i, u)
	// 		}
	// 	default:
	// 		fmt.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }

}
