package main

import (
	"fmt"
	"tbfo/dfa"
)

func main() {
	d, _ := dfa.ReadDFA("./assets/example.dfa.json")
	fmt.Println(d.Verify("a", "b"))
}
