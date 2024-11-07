package main

type Zone struct {
	Name  string
	Rooms map[Coordinates]Room
}

type Coordinates struct {
	X int
	Y int
}

const (
	North = 1
	East  = 2
	South = 3
	West  = 4
)

type Exit struct {
	Direction int
	Door      bool
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
