package exec

import (
	"bytes"
	"context"
	"io"
	"os/exec"
)

// RunCommand executes the given command with the given args returning both the standard output and error.
func RunCommand(
	ctx context.Context,
	name string,
	args ...string,
) (stdOut io.Reader, errOut io.Reader, err error) {
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	err = cmd.Run()
	stdOut = outBuf
	errOut = errBuf
	return
}
