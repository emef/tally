package lib

import (
	"time"
)

var epochMinute int32

func init() {
	epochMinute = calcNowEpochMinute()
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			<- ticker.C
			epochMinute = calcNowEpochMinute()
		}
	}()
}

func NowEpochMinute() int32 {
	return epochMinute
}

func calcNowEpochMinute() int32 {
	return int32(time.Now().Unix() / 60)
}
