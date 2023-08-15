package packagemanagers

import "os/exec"

type DpkgPackageManager struct{}

func (d DpkgPackageManager) Install(packageName, version string) error {
	// For dpkg, versioning isn't directly applied during installation.
	// Typically, the package file (e.g., .deb file) already contains the version.
	return exec.Command("dpkg", "-i", packageName).Run()
}

func (d DpkgPackageManager) Uninstall(packageName string) error {
	return exec.Command("dpkg", "-r", packageName).Run()
}

func (d DpkgPackageManager) List() error {
	return exec.Command("dpkg", "-l").Run()
}
