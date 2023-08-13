package packagemanagers

import (
	"fmt"
	"os/exec"
)

func UninstallSystemProgram(programName string) error {
	// Enclose program name in quotes
	programNameQuoted := fmt.Sprintf("'%s'", programName)

	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Get-AppxProvisionedPackage -Online | Where-Object { $_.DisplayName -eq %s } | Remove-AppxProvisionedPackage -Online", programNameQuoted))

	// Use CombinedOutput for capturing both stdout and stderr
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error with command:", err)
		fmt.Println("Command output:", string(out))
		return err
	}

	fmt.Println("Command output:", string(out))

	return nil
}
