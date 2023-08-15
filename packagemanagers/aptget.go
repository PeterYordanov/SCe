package packagemanagers

import (
	"fmt"
	"os/exec"
)

type AptGetPackageManager struct{}

func (a AptGetPackageManager) Install(packageName string, version string) error {
	fullPackage := fmt.Sprintf("%s=%s", packageName, version)
	return exec.Command("apt-get", "install", "-y", fullPackage).Run()
}

func (a AptGetPackageManager) Uninstall(packageName string) error {
	return exec.Command("apt-get", "remove", "-y", packageName).Run()
}

func (a AptGetPackageManager) List() error {
	return exec.Command("apt-get", "list", "--installed").Run()
}
