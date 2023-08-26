package packagemanagers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/PeterYordanov/SCe/core"
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

func (pm *ScoopPackageManager) List() ([]core.Package, error) {
	cmd := exec.Command("powershell", "-command", "(scoop list) | ForEach-Object { $_.Name + '-' + $_.Version }")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Failed to list packages: %s", err)
	}

	strOutput := strings.TrimSpace(string(output))

	packages := strings.Replace(strOutput, "Installed apps:", "", 1)
	packagesList := strings.Split(packages, "\r\n")

	result := make([]core.Package, 0)
	for _, value := range packagesList {
		tempSplit := strings.Split(value, "-")
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
