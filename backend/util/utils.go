package util

import (
	"fmt"
	"time"
)

func ModifyIntToTimeStr(intTime int64) string {
	currentUnixSec := time.Now().Unix()
	secSub := currentUnixSec - intTime

	switch {
	case 0 <= secSub && secSub < 60:
		return fmt.Sprintf("%d secs ago", secSub)
	case 60 <= secSub && secSub < 60*60:
		return fmt.Sprintf("%d mins ago", secSub/60)
	case 60*60 <= secSub && secSub < 60*60*24:
		return fmt.Sprintf("%d hrs ago", secSub/60/60)
	case 60*60*24 <= secSub:
		return fmt.Sprintf("%d days ago", secSub/60/60/24)
	default:
		return "0 secs ago"
	}
}

func ModifyBlockIntToTimeStr(intTime, pos int64) string {
	currentUnixSec := time.Now().Unix()
	fixPos := currentUnixSec - pos/1e9
	secSub := currentUnixSec - intTime/1e9 - fixPos + 1

	switch {
	case 0 <= secSub && secSub < 60:
		return fmt.Sprintf("%d secs ago", secSub)
	case 60 <= secSub && secSub < 60*60:
		return fmt.Sprintf("%d mins ago", secSub/60)
	case 60*60 <= secSub && secSub < 60*60*24:
		return fmt.Sprintf("%d hrs ago", secSub/60/60)
	case 60*60*24 <= secSub:
		return fmt.Sprintf("%d days ago", secSub/60/60/24)
	default:
		return "0 secs ago"
	}
}

//func ConvertSlotTimeToTimeStamp(soltTime int64) int64 {
//	t := consensus.Consensus.Timestamp{soltTime}
//	unixSec := t.ToUnixSec()
//
//	return unixSec
//}

func FormatUTCTime(intTime int64) string {
	t := time.Unix(intTime/1e9, intTime%1e9)
	return t.String()
}
