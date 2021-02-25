package mPlatformExecer

import "testing"

func Test_Exec(t *testing.T) {
	new(M_PlatformExecer).Exec("C:/Windows/System32/WindowsPowerShell/v1.0/powershell.exe", "mkdir", "hello")
}
