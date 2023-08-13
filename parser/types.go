package parser

type Package struct {
	Name           string `yaml:"name"`
	Version        string `yaml:"version"`
	PackageManager string `yaml:"package_manager"`
}

type Collection struct {
	Packages []Package `yaml:"packages"`
}
