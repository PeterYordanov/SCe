package packagemanagers

type AptGet struct{}

func (a AptGet) Install(packageName string, version string) error {
	fullPackage := fmt.Sprintf("%s=%s", packageName, version)
	return exec.Command("apt-get", "install", "-y", fullPackage).Run()
}

func (a AptGet) Uninstall(packageName string) error {
	return exec.Command("apt-get", "remove", "-y", packageName).Run()
}

func (a AptGet) List() error {
	return exec.Command("apt-get", "list", "--installed").Run()
}
