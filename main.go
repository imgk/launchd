package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if os.Args[1] == "start" {
		if err := runCmd(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	}
}

func runCmd(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
