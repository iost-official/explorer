package model

import (
	"fmt"
	"time"

	"github.com/iost-official/prototype/consensus/common"
)

func modifyIntToTimeStr(intTime int64) string {
	currentUnixSec := time.Now().Unix()
	secSub := currentUnixSec - intTime

	switch {
	case 0 <= secSub && secSub < 60:
		return fmt.Sprintf("%d secs ago", secSub)
	case 60 <= secSub && secSub < 60 * 60:
		return fmt.Sprintf("%d mins ago", secSub / 60)
	case 60 * 60 <= secSub && secSub < 60 * 60 * 24:
		return fmt.Sprintf("%d hrs ago", secSub / 60 / 60)
	case 60 * 60 * 24 <= secSub:
		return fmt.Sprintf("%d days ago", secSub / 60 / 60 / 24)
	default:
		return "0 secs ago"
	}
}

func modifyBlockIntToTimeStr(intTime int64) string {
	currentUnixSec := time.Now().Unix()
	secSub := currentUnixSec - intTime

	if secSub - 12 > 0 {
		secSub = secSub - 12
	} else if secSub - 9 > 0 {
		secSub = secSub - 9
	} else if secSub - 6 > 0 {
		secSub = secSub - 6
	} else if secSub - 3 > 0 {
		secSub = secSub - 3
	}

	switch {
	case 0 <= secSub && secSub < 60:
		return fmt.Sprintf("%d secs ago", secSub)
	case 60 <= secSub && secSub < 60 * 60:
		return fmt.Sprintf("%d mins ago", secSub / 60)
	case 60 * 60 <= secSub && secSub < 60 * 60 * 24:
		return fmt.Sprintf("%d hrs ago", secSub / 60 / 60)
	case 60 * 60 * 24 <= secSub:
		return fmt.Sprintf("%d days ago", secSub / 60 / 60 / 24)
	default:
		return "0 secs ago"
	}
}

func ConvertSlotTimeToTimeStamp(soltTime int64) int64 {
	t := consensus_common.Timestamp{soltTime}
	unixSec := t.ToUnixSec()

	return unixSec
}

func formatUTCTime(intTime int64) string {
	t := time.Unix(intTime, 0)
	return t.String()
}
