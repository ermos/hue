package logger

import "log"

// 0 = No Alert
// 1 = Info
// 2 = Error
// 3 = All
var Level = 0

func SetLevel(lvl int) {
	if lvl > 3 || lvl < 0 {
		log.Fatalf("debug mode %d is not available", lvl)
	}
	Level = lvl
}
