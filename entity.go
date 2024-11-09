package main

type Stat struct {
	Current int
	Max     int
}

type Character struct {
	Name   string
	Class  string
	HP     Stat
	MP     Stat
	MVP    Stat
	TNL    Stat
	AB     int
	AC     int
	Skills map[string]Stat
	STR    int
	DEX    int
	INT    int
	WIS    int
	CON    int
	LUK    int
	Flags  []string
	Tags   []string
}

func CreateCharacter(name string, class string, str int, dex int, con int, intel int, wis int, luk int) Character {
	return Character{
		Name:  name,
		Class: class,
		STR:   str,
		DEX:   dex,
		INT:   intel,
		WIS:   wis,
		LUK:   luk,
		CON:   con,
		HP: Stat{
			Current: con * 10,
			Max:     con * 10,
		},
		MP: Stat{
			Current: (con + wis) * 10,
			Max:     (con + wis) * 10,
		},
		MVP: Stat{
			Current: dex * 355,
			Max:     dex * 355,
		},
		TNL: Stat{
			Current: 0,
			Max:     1000,
		},
		AB: 1,
	}
}
