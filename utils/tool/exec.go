package tool

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// 执行终端命令
func ExecCmd(dir string, name string, args ...string) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ()
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return stdout.String(), fmt.Errorf("%s\n%s", err.Error(), stderr.String())
	}
	return stdout.String(), nil
}
