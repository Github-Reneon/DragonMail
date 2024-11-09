package main

import "math/rand"

type Zone struct {
	Name   string
	Rooms  map[Coordinates]Room
	Random *rand.Rand
}

type Coordinates struct {
	X int
	Y int
}

type Exit struct {
	Direction int
	Door      bool
	Opened    bool
	Locked    bool
	LockDC    int
}

type Room struct {
	Name        string
	Description string
	Characters  []Character
	Objects     string
	Exits       []Exit
}
