package main

import (
	"fmt"
)

func main() {

	m := make(map[string]interface{})
	m1 := make(map[string]interface{})
	m1["ip"] = "127.0.0.1"
	m["127.0.0.1"] = m1
	m["127.0.0.2"] = m1

	var agents []string
	for s, v := range m {
		vmap, ok := v.(map[string]interface{})
		if ok {
			if vmap["ip"] == s {
				agents = append(agents, s)
			}
		}

	}
	fmt.Println(agents)
}
