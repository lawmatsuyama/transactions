package domain

import "time"

// Now is a type func to return the current time. It's global for test purposes
var Now = func() time.Time {
	return time.Now()
}

func TimeSaoPaulo(t time.Time) time.Time {
	return t.In(time.FixedZone("America/Sao_Paulo", -3*60*60))
}
