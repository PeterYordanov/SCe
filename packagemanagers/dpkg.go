package packagemanagers

type Dpkg struct{}

func (d Dpkg) Install(packageName, version string) error {
	// For dpkg, versioning isn't directly applied during installation.
	// Typically, the package file (e.g., .deb file) already contains the version.
	return exec.Command("dpkg", "-i", packageName).Run()
}

func (d Dpkg) Uninstall(packageName string) error {
	return exec.Command("dpkg", "-r", packageName).Run()
}

func (d Dpkg) List() error {
	return exec.Command("dpkg", "-l").Run()
}
