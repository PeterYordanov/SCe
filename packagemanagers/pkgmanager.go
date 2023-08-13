package packagemanagers

type PackageManager interface {
	Install(packageName, version string) error
	Uninstall(packageName string) error
	List() ([]string, error)
}
