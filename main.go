package main

func main() {
	functionMap := make(map[string]func())

	functionMap["look"] = look

	functionMap["look"]()

	if function, exists := functionMap["looks"]; exists {
		function()
	} else {
		badinput()
	}
}