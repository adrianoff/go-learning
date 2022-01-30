package main

import "fmt"

type Terminator struct {
	On          bool
	Ammo, Power int
}

func (terminator *Terminator) Shoot() bool {
	if !terminator.On {
		return false
	}
	if terminator.Ammo > 0 {
		terminator.Ammo--
		return true
	}

	return false
}

func (terminator *Terminator) RideBike() bool {
	if !terminator.On {
		return false
	}
	if terminator.Power > 0 {
		terminator.Power--
		return true
	}

	return false
}

func main() {
	term := *new(Terminator)

	fmt.Print(term)
}
