package main

import "syscall"
import "os"
import "os/exec"

func main() {

	// Not working on Windows 10
	binary, lookErr := exec.LookPath("ver")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{""}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
