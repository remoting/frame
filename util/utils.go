package util

import (
	"bytes"
	"fmt"
	"reflect"
	"time"

	"golang.org/x/crypto/ssh"
)

func TimeWatcher(name string) {
	start := time.Now()
	defer func() {
		cost := time.Since(start)
		fmt.Printf("%s: %v\n", name, cost)
	}()
}

func RunCommand(client *ssh.Client, command string) (stdout string, err error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var buf bytes.Buffer
	session.Stdout = &buf
	err = session.Run(command)
	if err != nil {
		return "", err
	}
	stdout = buf.String()
	return
}
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	} else {
		return i == nil
	}
}
