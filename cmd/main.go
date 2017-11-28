package main

import (
	"fmt"

	"github.com/Rufio0425/tool"
)

func main() {
	var myname string
	var toolresponse string

	nameQuestion := "What is your name?"
	fmt.Println(nameQuestion)
	fmt.Scanf("%s", &myname)

	if tool.IsTool([]byte(myname)) {
		toolresponse = "You're a tool... sorry bud."
	} else {
		toolresponse = "Congrats, you're not a tool. Celebrate!"
	}

	fmt.Printf("Hello %v! %v\n", myname, toolresponse)
}
