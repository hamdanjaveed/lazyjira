package main

import (
	"context"
	"fmt"
	"log"

	lgerrors "github.com/hamdanjaveed/lazygit/pkg/errors"
	lgexec "github.com/hamdanjaveed/lazygit/pkg/exec"
)

func run() error {
	cout, cerr, err := lgexec.RunCommand(
		context.Background(),
		"jira",
		[]string{
			"issues",
			"list",
			"--plain",
			"--no-truncate",
		}...,
	)
	if err != nil {
		return lgerrors.Wrapf(err, "failed to run jira")
	}

	fmt.Println(cout, cerr)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
