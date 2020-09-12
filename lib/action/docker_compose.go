package action

import (
	"fmt"
	"os/exec"

	"github.com/rapde/rap/lib/utils"
)

// ExecDockerCompose 执行 docker-compose 命令
func ExecDockerCompose(args ...string) (output string, err error) {
	dcBin, _ := utils.GetBinPath("docker-compose")

	out, err := exec.Command(dcBin, args...).Output()
	if err != nil {
		return "", fmt.Errorf("Failed to execute command: %s, %s", args, string(out))
	}

	return string(out), nil
}
