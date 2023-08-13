package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func (p *Runbook) Parse() (string, error) {

	// Parse Runbook
	data, err := os.ReadFile(p.RunbookFilePath)

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	yamlParser := NewYamlWrapper[Runbook]()

	err = yamlParser.Parse(string(data))

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	// Read includes and parse them

	//collections := yamlParser.data.Runbook
	includes := yamlParser.data.Include

	runs := make([]Package, 0)

	for _, value := range includes {
		data, err := os.ReadFile(filepath.Join(p.RunbookDirPath, value))

		if err != nil {
			fmt.Print(err)
			return "", err
		}

		yamlParserIncludes := NewYamlWrapper[Collection]()

		err = yamlParserIncludes.Parse(string(data))

		if err != nil {
			fmt.Print(err)
			return "", err
		}

		runs = append(runs, yamlParserIncludes.data.Packages...)
	}

	p.Runs = append(p.Runs, runs...)

	return string(data), err
}

func (p Runbook) Run() (string, error) {
	fmt.Println(p.Runs)

	for _, value := range p.Runs {
		fmt.Println(value)

		switch strings.ToLower(value.PackageManager) {
		case "choco":
			fmt.Println("PackageManager is choco")
		case "scoop":
			fmt.Println("PackageManager is scoop")
		case "apt-get":
			fmt.Println("PackageManager is apt-get")
		case "winget":
			fmt.Println("PackageManager is winget")
		case "snap":
			fmt.Println("PackageManager is snap")
		default:
			return "", fmt.Errorf("Unknown package manager: %s", value.PackageManager)
		}
	}

	return "", nil
}
