package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PeterYordanov/SCe/packagemanagers"
)

type Runbook struct {
	Include         []string `yaml:"includes"`
	Runbook         []string `yaml:"run"`
	RunbookFilePath string
	RunbookDirPath  string
	Runs            []Package
}

func NewRunbook(runbookPath string) *Runbook {
	return &Runbook{
		RunbookFilePath: runbookPath,
		RunbookDirPath:  filepath.Dir(runbookPath),
	}
}

// TODO: Move and rename
func stringExists(target string, slice []string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func (p *Runbook) Parse() error {

	// Parse Runbook
	data, err := os.ReadFile(p.RunbookFilePath)

	if err != nil {
		return err
	}

	yamlParser := NewYamlWrapper[Runbook]()

	err = yamlParser.Parse(string(data))

	if err != nil {
		return err
	}

	// Read includes and parse them
	for _, value := range yamlParser.data.Include {
		includePath := filepath.Join(p.RunbookDirPath, value)
		_, err := os.ReadFile(includePath)
		if err != nil {
			return err
		}

		fmt.Println("Found include file:", includePath)
	}

	runs := make([]Package, 0)
	for _, value := range yamlParser.data.Runbook {
		// If the run is in the include list
		if stringExists(value, yamlParser.data.Include) {
			data, err := os.ReadFile(filepath.Join(p.RunbookDirPath, value))

			if err != nil {
				return err
			}

			yamlParserIncludes := NewYamlWrapper[Collection]()

			err = yamlParserIncludes.Parse(string(data))

			if err != nil {
				return err
			}

			runs = append(runs, yamlParserIncludes.data.Packages...)
		} else {
			return fmt.Errorf("Unexpected collection in runbook: %s", value)
		}
	}

	p.Runs = append(p.Runs, runs...)

	return err
}

func (p Runbook) Run() (string, error) {
	for _, value := range p.Runs {
		switch strings.ToLower(value.PackageManager) {
		case "choco":
			fmt.Println("PackageManager is choco")
			pkgManager := packagemanagers.NewChocolatey()
			pkgManager.Install(value.Name, value.Version)
		case "scoop":
			fmt.Println("PackageManager is scoop")
			pkgManager := packagemanagers.NewScoop()
			pkgManager.Install(value.Name, value.Version)
		case "apt-get":
			fmt.Println("PackageManager is apt-get")
		case "snap":
			fmt.Println("PackageManager is snap")
		default:
			return "", fmt.Errorf("Unknown package manager: %s", value.PackageManager)
		}
	}

	return "", nil
}
