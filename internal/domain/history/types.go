package history

import "time"

type History struct {
	ShinobiID int

	DateAvailable *time.Time
	DateReserved  *time.Time
	DateOnMission *time.Time
	DateInjured   *time.Time
	DateOnLeave   *time.Time
	DateDisabled  *time.Time
	DateDead      *time.Time
	DateNukenin   *time.Time
}
