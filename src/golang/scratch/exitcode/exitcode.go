package main

import (
	"os"
	"os/exec"
)

// getExitCode passes along the return value from a command run by os/exec
func getExitCode(name string, arg ...string) int {
	cmd := exec.Command(name, arg...)
	_ = cmd.Run()
	return cmd.ProcessState.ExitCode()
	panic("something wrong!")
}

func main() {
	os.Exit(getExitCode("sh", "-c", "false"))
}
