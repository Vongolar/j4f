package mPlatformExecer

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type M_PlatformExecer struct{}

func (m *M_PlatformExecer) Exec(name string, arg ...string) {
	cmd := exec.CommandContext(context.Background(), name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
