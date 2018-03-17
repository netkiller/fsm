[![wercker status](https://app.wercker.com/status/517d98fe7a8da9bf9a6060e7906c0d17/s "wercker status")](https://app.wercker.com/project/bykey/517d98fe7a8da9bf9a6060e7906c0d17)
[![Coverage Status](https://img.shields.io/coveralls/looplab/state.svg)](https://coveralls.io/r/looplab/state)
[![GoDoc](https://godoc.org/github.com/looplab/state?status.svg)](https://godoc.org/github.com/looplab/state)
[![Go Report Card](https://goreportcard.com/badge/looplab/state)](https://goreportcard.com/report/looplab/state)


# State for Go

State is a finite state machine for Go.

It is heavily based on two State implementations:

- Javascript Finite State Machine, https://github.com/jakesgordon/javascript-state-machine

- Fysom for Python, https://github.com/oxplot/fysom (forked at https://github.com/mriehl/fysom)

For API docs and examples see http://godoc.org/github.com/looplab/state


# Basic Example

From examples/simple.go:

```go
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
```


# Usage as a struct field

From examples/struct.go:

```go
package main

import (
    "fmt"
    "github.com/looplab/state"
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
```


# License

State is licensed under Apache License 2.0

http://www.apache.org/licenses/LICENSE-2.0
