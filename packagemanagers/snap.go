package packagemanagers

type Snap struct{}

func (s Snap) Install(packageName string, version string) error {
	// For snap, you specify channels instead of versions directly
	channel := fmt.Sprintf("--channel=%s", version)
	return exec.Command("snap", "install", packageName, channel).Run()
}

func (s Snap) Uninstall(packageName string) error {
	return exec.Command("snap", "remove", packageName).Run()
}

func (s Snap) List() error {
	return exec.Command("snap", "list").Run()
}
