package packagemanagers

import (
	"fmt"
	"os/exec"
)

type SnapPackageManager struct{}

func (s SnapPackageManager) Install(packageName string, version string) error {
	// For snap, you specify channels instead of versions directly
	channel := fmt.Sprintf("--channel=%s", version)
	return exec.Command("snap", "install", packageName, channel).Run()
}

func (s SnapPackageManager) Uninstall(packageName string) error {
	return exec.Command("snap", "remove", packageName).Run()
}

func (s SnapPackageManager) List() error {
	return exec.Command("snap", "list").Run()
}
