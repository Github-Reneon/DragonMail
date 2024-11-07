package main

import "fmt"

func main() {
	
	functionMap := map[string]func(){
		"look": look,
		"l": look,
		"attack": attack,
		"a": attack,
		"exits": exits,
		"e": exits,
		"quit": quit,
		"qq": quit,
	}

	for {
		if function, exists := functionMap["looks"]; exists {
			function()
		} else {
			badinput()
		}
	}
}