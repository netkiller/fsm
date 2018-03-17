// +build ignore

package main

import (
	"fmt"
	"github.com/netkiller/state"
)

func main() {
	state := state.NewState(
		"idle",
		state.Events{
			{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
			{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
			{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
		},
		state.Callbacks{
			"scan": func(e *state.Event) {
				fmt.Println("after_scan: " + e.State.Current())
			},
			"working": func(e *state.Event) {
				fmt.Println("working: " + e.State.Current())
			},
			"situation": func(e *state.Event) {
				fmt.Println("situation: " + e.State.Current())
			},
			"finish": func(e *state.Event) {
				fmt.Println("finish: " + e.State.Current())
			},
		},
	)

	fmt.Println(state.Current())

	err := state.Event("scan")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("1:" + state.Current())

	err = state.Event("working")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2:" + state.Current())

	err = state.Event("situation")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("3:" + state.Current())

	err = state.Event("finish")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("4:" + state.Current())

}
