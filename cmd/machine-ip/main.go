package main

import (
	"errors"
	"fmt"
	"github.com/fnitschmann/machine-ip"
	"os"
	"strings"
)

var (
	// Should just be used to simply end the process
	ErrUsage = errors.New("usage")

	// ErrUnknownCommand is returned when a CLI command is not specified
	ErrUnknownCommand = errors.New("command unknown")
)

func main() {
	r := Run(os.Args[1:]...)
	if r != nil && r == ErrUsage {
		os.Exit(2)
	} else if r != nil {
		fmt.Println(r.Error())
		os.Exit(1)
	}
}

func Run(args ...string) error {
	if len(args) == 0 {
		fmt.Printf(Usage())
		return ErrUsage
	} else {
		switch args[0] {
		case "local":
			return newLocalCommand()
		case "public":
			return newPublicCommand()
		default:
			return ErrUnknownCommand
		}
	}
}

func Usage() string {
	return strings.TrimLeft(`
machine-ip is a tool to get the machines' IPv4 addresses

Usage:
	machine-ip COMMAND

The commands are:

local	get local network IPv4 address of the machine
public	get public (internet) IPv4 address of the machine (uses ipify.org service)
	`, "\n")
}

func newLocalCommand() error {
	ip, err := ip.GetLocalMachineIp()
	if err != nil {
		return err
	}

	fmt.Println(ip)
	return nil
}

func newPublicCommand() error {
	ip, err := ip.GetPublicMachineIp()
	if err != nil {
		return err
	}

	fmt.Println(ip)
	return nil
}
