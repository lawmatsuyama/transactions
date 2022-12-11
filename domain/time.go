package domain

import "time"

var Now = func() time.Time {
	return time.Now()
}

func TimeSaoPaulo(t time.Time) time.Time {
	return t.In(time.FixedZone("America/Sao_Paulo", -3*60*60))
}
