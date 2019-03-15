package blockchain

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	RPCAddress  = "127.0.0.1:30002"
	PasswordKey = []byte("temp")
	BPRegisterPassword = "temp"
)

func readPassword(prompt string) (pw []byte, err error) {
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

func InitConfig() {
	RPCAddress = viper.GetString("rpcHost")
	BPRegisterPassword = viper.GetString("BPRegisterPassword")
	fmt.Println("RPCHost:", RPCAddress)
	bytePassword, err := readPassword("Enter Password:  ")
	if err != nil {
		fmt.Println(err)
	}
	PasswordKey = bytePassword
}
