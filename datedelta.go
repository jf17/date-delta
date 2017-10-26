package datedelta

import "time"

func GetMinute(d time.Time) int64 {
	t := time.Now()

	nowTime := t.Unix()
	lastTime := d.Unix()

	deltaMin := (lastTime -nowTime) / 60

	return deltaMin
}
