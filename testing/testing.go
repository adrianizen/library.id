package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"adrianizen/library.id/internal/config"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")

	config.RootDirectory = dir

	// get `go` executable path
	goExecutable, _ := exec.LookPath("go")

	// construct `go version` command
	cmdGoVer := &exec.Cmd{
		Path:   goExecutable,
		Args:   []string{goExecutable, "test", "-p", "1", "-v", "./../..."},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// run `go version` command
	if err := cmdGoVer.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
