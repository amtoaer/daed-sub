package cmd

import (
	"fmt"
	"os"
)

func Execute() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("invalid arguments")
	}
	switch os.Args[1] {
	case "handle":
		return Handle()
	case "daemon":
		return Daemon()
	case "init":
		return Init()
	}
	return nil
}
