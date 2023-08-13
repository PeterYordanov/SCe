package packagemanagers

import (
	"fmt"
	"os/exec"
	"strings"
)

type ScoopPackageManager struct{}

func NewScoop() *ScoopPackageManager {
	return &ScoopPackageManager{}
}

func (pm *ScoopPackageManager) Install(packageName, version string) error {
	fmt.Println(fmt.Sprintf("Installing %s with version %s", packageName, version))

	cmdStr := fmt.Sprintf("scoop install %s@%s", packageName, version)
	cmd := exec.Command("cmd", "/c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to uninstall package: %s", string(output))
		return fmt.Errorf("Failed to install package: %s", err)
	}

	fmt.Println(string(output))
	return nil
}

func (pm *ScoopPackageManager) Uninstall(packageName string) error {
	fmt.Println(fmt.Sprintf("Uninstalling %s", packageName))

	cmdStr := fmt.Sprintf("scoop uninstall %s", packageName)
	cmd := exec.Command("cmd", "/c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to uninstall package: %s", string(output))
		return fmt.Errorf("Failed to uninstall package: %s", err)
	}

	fmt.Println(string(output))
	return nil
}

func (pm *ScoopPackageManager) List() ([]string, error) {
	cmd := exec.Command("scoop", "list")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Failed to list packages: %s", err)
	}

	packages := strings.Split(strings.TrimSpace(string(output)), "\r\n")
	return packages, nil
}
