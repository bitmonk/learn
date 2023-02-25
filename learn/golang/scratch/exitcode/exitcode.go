package exitcode

import (
	"os/exec"
)

// GetExitCode passes along the return value from a command run by os/exec
func GetExitCode(name string, arg ...string) int {
	cmd := exec.Command(name, arg...)
	_ = cmd.Run()
	return cmd.ProcessState.ExitCode()
}
