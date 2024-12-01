package types

import "time"

type Payment struct {
	ID     uint64
	Label  string
	Amount float64
	Time   time.Time
}
