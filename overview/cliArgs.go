package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var JSONrecord = `{
	"Flag": true,
	"Array": ["a","b","c"],
	"Entity": {
		"a1": "b1",
		"a2": "b2",
		"Value": -456,
		"Null": null
		},
	"Message": "Hello Go!"
	}`

func TypeSwitch(m map[string]interface{}) {
	fmt.Println("**************************")
	for k, v := range m {
		fmt.Println("KEY: ", k)
		switch c := v.(type) {
		// c = {"a": "chu","b": "dep","c": "trai"]
		case string:
			fmt.Println("Is a string: ", k, c)
		case float64:
			fmt.Println("Is a float64: ", k, c)
		case int:
			fmt.Println("Is an int: ", k, c)
		case map[string]any:
			fmt.Println("Is a map!", k, c)
			//TypeSwitch(c)
			TypeSwitch(v.(map[string]any))
		default:
			fmt.Printf("... Is %v: %T\n", v, c)
		}
	}
	return
}
func ExploreMap(m map[string]interface{}) {
	fmt.Println("-------------------------")
	for k, v := range m {
		embMap, ok := v.(map[string]interface{})
		if ok {
			fmt.Println("clt")
			fmt.Printf("{\"%v\": \n", k)
			ExploreMap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v, %v\n", k, v)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default json records")
	} else {
		JSONrecord = os.Args[1]
	}
	JSONMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	//ExploreMap(JSONMap)
	TypeSwitch(JSONMap)
}
