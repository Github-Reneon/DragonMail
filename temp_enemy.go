package main

import (
	"dragonmail/ansi"

	"github.com/Github-Reneon/dicetools"
)

func CreateBat() Character {
	health := 100 + (dicetools.RollNotation("1d4") * 10)
	return Character{
		Name:  "Bat",
		Class: "DefaultEnemy",
		STR:   3,
		DEX:   5,
		INT:   0,
		WIS:   1,
		LUK:   1,
		CON:   10,
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
		Tags: []string{
			ansi.Blue + "(Flying)" + ansi.Reset,
		},
	}
}
