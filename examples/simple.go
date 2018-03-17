// +build ignore

package main

import (
	"fmt"
	"github.com/netkiller/state"
)

func main() {
	state := state.NewState(
		"closed",
		state.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		state.Callbacks{},
	)

	fmt.Println(state.Current())

	err := state.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(state.Current())

	err = state.Event("close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(state.Current())
}
