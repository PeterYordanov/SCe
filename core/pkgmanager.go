package core

type PackageManager interface {
	Install(packageName, version string) error
	Uninstall(packageName string) error
	List() ([]Package, error)
}
