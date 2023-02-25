package main

import (
	"os"
	"learn/golang/scratch/exitcode"
)

func main() {
	os.Exit(exitcode.GetExitCode("sh", "-c", "false"))
}
