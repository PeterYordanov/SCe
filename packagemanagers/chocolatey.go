package packagemanagers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/PeterYordanov/SCe/core"
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

func (pm *ChocolateyPackageManager) List() ([]core.Package, error) {
	cmd := exec.Command("powershell", "-nologo", "-noprofile", `(choco list --local-only) | Select -SkipLast 1 | ForEach-Object { $_.Split(" ")[0] + "|" + $_.Split(" ")[1] }`)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Failed to list packages: %s", err)
	}

	packages := strings.Split(strings.TrimSpace(string(output)), "\r\n")

	result := make([]core.Package, 0)
	for _, value := range packages {
		tempSplit := strings.Split(value, "|")
		packageName := tempSplit[0]
		packageVersion := tempSplit[1]

		result = append(result, core.Package{
			PackageManager: "Chocolatey",
			Name:           packageName,
			Version:        packageVersion,
		})
	}

	return result, nil
}
