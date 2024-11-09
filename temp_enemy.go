package main

import (
	"dragonmail/ansi"

	"math/rand"
	"time"

	"github.com/Github-Reneon/dicetools"
)

func CreateBat(tag string, rand *rand.Rand) Character {
	health := 100 + (dicetools.RollNotation("1d4", rand) * 10)
	time.Sleep(1 * time.Millisecond)
	return Character{
		Name:  "Bat",
		Class: "DefaultEnemy",
		Tags: []string{
			"bat",
		},
		STR: 3,
		DEX: 5,
		INT: 0,
		WIS: 1,
		LUK: 1,
		CON: 10,
		HP: Stat{
			Current: health,
			Max:     health,
		},
		MP: Stat{
			Current: 0,
			Max:     0,
		},
		MVP: Stat{
			Current: 0,
			Max:     0,
		},
		TNL: Stat{
			Current: 0,
			Max:     10,
		},
		AB: 1,
		AC: 10,
		Flags: []string{
			ansi.Blue + "(" + tag + ")" + ansi.Reset,
		},
	}
}
