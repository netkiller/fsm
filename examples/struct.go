// +build ignore

package main

import (
	"fmt"
	"github.com/netkiller/state"
)

type Door struct {
	To  string
	State *state.State
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.State = state.NewState(
		"closed",
		state.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		state.Callbacks{
			"enter_state": func(e *state.Event) { d.enterState(e) },
		},
	)

	return d
}

func (d *Door) enterState(e *state.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func main() {
	door := NewDoor("heaven")

	err := door.State.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.State.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}
