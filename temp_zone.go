package main

import "dragonmail/directions"

func createZone() Zone {
	return Zone{
		Name: "The Dark Cave",
		Rooms: map[Coordinates]Room{
			{
				X: 0,
				Y: 0,
			}: {
				Name: "A dark cave",
				Description: "The smell is dank and repulsive.\nThere is a trickling stream in this dark cave that flows to the east.There's a constant background noise of quiet but animalistic screeching.",
				Characters: []Character{
					CreateBat("sleeping"),
				},
				Objects: "Spoon",
				Exits: []Exit{
					{
						Direction: directions.East,
						Door:      false,
					},
				},
			},
			{
				X: 1,
				Y: 0,
			}: {
				Name: "Deeper into the cave",
				Description: "You see that the screeching is louder!\nA bat's nest is before you!",
				Characters: []Character{
					CreateBat("flying"),
					CreateBat("sleeping"),
				},
				Objects: "Gold!",
				Exits: []Exit{
					{
						Direction: directions.West,
						Door:      false,
					},
				},
			},
		},
	}
}
