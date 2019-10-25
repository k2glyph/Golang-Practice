package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
		break
	default:
		panic("What the fuck??")
	}
}

func run() {
	fmt.Printf("running %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Cloneflags:   syscall.RTF_CLONING,
	// 	Unshareflags: syscall.CLONE_NEWNS,
	// }
	must(cmd.Run())
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
