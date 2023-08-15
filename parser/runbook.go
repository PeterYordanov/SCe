package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/PeterYordanov/SCe/core"
	"github.com/PeterYordanov/SCe/packagemanagers"
)

type Runbook struct {
	Include         []string `yaml:"includes"`
	Runbook         []string `yaml:"run"`
	RunbookFilePath string
	RunbookDirPath  string
	Runs            []core.Package
}

func NewRunbook(runbookPath string) *Runbook {
	return &Runbook{
		RunbookFilePath: runbookPath,
		RunbookDirPath:  filepath.Dir(runbookPath),
	}
}

func (p *Runbook) Parse() error {
	yamlParser := NewYamlWrapper[Runbook]()
	parseIncludesError := yamlParser.ReadAndParse(p.RunbookFilePath)

	if parseIncludesError != nil {
		return parseIncludesError
	}

	// Validate Includes
	for _, value := range yamlParser.data.Include {
		core.ValidateIncludeExists(filepath.Join(p.RunbookDirPath, value))
	}

	// Validate runs
	runs := make([]core.Package, 0)
	for _, value := range yamlParser.data.Runbook {

		if core.CollectionExistsInIncludes(value, yamlParser.data.Include) {
			yamlParserIncludes := NewYamlWrapper[core.Collection]()
			parseCollectionErr := yamlParserIncludes.ReadAndParse(filepath.Join(p.RunbookDirPath, value))

			if parseCollectionErr != nil {
				return parseCollectionErr
			}

			runs = append(runs, yamlParserIncludes.data.Packages...)
		} else {
			return fmt.Errorf("Unexpected collection in runbook: %s", value)
		}

	}

	p.Runs = append(p.Runs, runs...)

	return nil
}

func (p Runbook) Run() error {
	for _, value := range p.Runs {
		switch strings.ToLower(value.PackageManager) {
		case "choco":
			fmt.Println("PackageManager is choco")
			pkgManager := packagemanagers.NewChocolatey()

			fmt.Println(pkgManager.List())

			isInstalled, err := core.IsPackageInstalled(pkgManager, value.Name, value.Version)

			if err != nil {
				return err
			}

			if isInstalled {
				fmt.Printf("Package '%s-%s' is already installed, skipping...\n", value.Name, value.Version)
			} else {
				installErr := pkgManager.Install(value.Name, value.Version)
				if installErr != nil {
					return installErr
				}
			}
		case "scoop":
			fmt.Println("PackageManager is scoop")
			pkgManager := packagemanagers.NewScoop()
			installErr := pkgManager.Install(value.Name, value.Version)
			if installErr != nil {
				return installErr
			}
		case "apt-get":
			fmt.Println("PackageManager is apt-get")
		case "snap":
			fmt.Println("PackageManager is snap")
		default:
			return fmt.Errorf("Unknown package manager: %s", value.PackageManager)
		}
	}

	return nil
}
