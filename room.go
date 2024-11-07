package main

type Zone struct {
	Name  string
	Rooms map[Coordinates]Room
}

type Coordinates struct {
	X int
	Y int
}

type Room struct {
	Name        string
	Description string
	Characters  []Character
	Objects     string
}
