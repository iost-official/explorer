package util

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"time"
)

// ReadPassword ...
func ReadPassword(prompt string) (pw []byte, err error) {
	fd := int(os.Stdin.Fd())
	if terminal.IsTerminal(fd) {
		fmt.Fprint(os.Stderr, prompt)
		pw, err = terminal.ReadPassword(fd)
		fmt.Fprintln(os.Stderr)
		return
	}

	var b [1]byte
	for {
		n, err := os.Stdin.Read(b[:])
		// terminal.ReadPassword discards any '\r', so we do the same
		if n > 0 && b[0] != '\r' {
			if b[0] == '\n' {
				return pw, nil
			}
			pw = append(pw, b[0])
			// limit size, so that a wrong input won't fill up the memory
			if len(pw) > 1024 {
				err = errors.New("password too long")
			}
		}
		if err != nil {
			// terminal.ReadPassword accepts EOF-terminated passwords
			// if non-empty, so we do the same
			if err == io.EOF && len(pw) > 0 {
				err = nil
			}
			return pw, err
		}
	}
}

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
