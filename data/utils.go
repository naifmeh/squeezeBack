package data

import "time"

func CompareAuthorizationTime(beg, end int64) bool {
	timeNow := time.Now().Unix()
	if beg > timeNow {
		return false
	} else if timeNow > end {
		return false
	}
	return true
}

