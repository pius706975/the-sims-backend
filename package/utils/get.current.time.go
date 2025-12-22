package utils

import "time"

var makassarLoc *time.Location

func init() {
	loc, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		makassarLoc = time.FixedZone("WITA", 8*60*60)
		return
	}
	makassarLoc = loc
}

func GetCurrentTime() time.Time {
	return time.Now().In(makassarLoc)
}
