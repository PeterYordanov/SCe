package packagemanagers

import (
	"fmt"
	"os/exec"
	"strings"
)

type ChocolateyPackageManager struct{}

func NewChocolatey() *ChocolateyPackageManager {
	return &ChocolateyPackageManager{}
}

func (pm *ChocolateyPackageManager) Install(packageName, version string) error {
	fmt.Println(fmt.Sprintf("Installing %s with version %s", packageName, version))

	cmdStr := fmt.Sprintf("choco install %s -y --version %s", packageName, version)
	cmd := exec.Command("cmd", "/c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to uninstall package: %s", string(output))
		return fmt.Errorf("Failed to install package: %s", err)
	}

	fmt.Println(string(output))
	return nil
}

func (pm *ChocolateyPackageManager) Uninstall(packageName string) error {
	fmt.Println(fmt.Sprintf("Uninstalling %s", packageName))

	cmdStr := fmt.Sprintf("choco uninstall %s -y", packageName)
	cmd := exec.Command("cmd", "/c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to uninstall package: %s", string(output))
		return fmt.Errorf("Failed to uninstall package: %s", err)
	}

	fmt.Println(string(output))
	return nil
}

func (pm *ChocolateyPackageManager) List() ([]string, error) {
	cmd := exec.Command("choco", "list")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Failed to list packages: %s", err)
	}

	packages := strings.Split(strings.TrimSpace(string(output)), "\r\n")
	return packages, nil
}
