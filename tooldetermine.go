package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var myname string
	var toolresponse string

	nameQuestion := "What is your name?"
	fmt.Println(nameQuestion)
	fmt.Scanf("%s", &myname)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomNum := r1.Intn(100)

	if randomNum > 51 {
		toolresponse = "You're a tool... sorry bud"
	} else {
		toolresponse = "Congrats, you're not a tool. Celebrate"
	}

	fmt.Print("Hello ", myname+" "+toolresponse)
	fmt.Println()
}
