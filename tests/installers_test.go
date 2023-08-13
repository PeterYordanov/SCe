package tests

// go test ./tests

import (
	"fmt"
	"testing"

	"github.com/PeterYordanov/SCe/packagemanagers"
	"github.com/PeterYordanov/SCe/parser"
)

func TestRunbook(t *testing.T) {
	// Arrange: Create a new runbook
	runbook := parser.NewRunbook("C:\\Projects\\SCe\\sample\\sample-runbook.yml")

	// Act: Parse and run the runbook
	runbook.Parse()
	runbook.Run()

	// Assert: Add assertions if needed
}

func TestPackageManagerInstaller(t *testing.T) {
	// Arrange: Create a new package manager installer
	installer := packagemanagers.NewPackageManagerInstaller(packagemanagers.Chocolatey)

	// Act: Install the package manager
	err := installer.Install()

	// Assert: Check for errors and expectations
	if err != nil {
		t.Errorf("Error installing package manager: %v", err)
	}

	// Act: Get the version of the package manager
	version, err := installer.GetVersion()

	// Assert: Check for errors and expectations
	if err != nil {
		t.Errorf("Error getting package manager version: %v", err)
	}
	fmt.Println(version) // Consider logging instead of printing during tests

	// Act: Check if the package manager is installed
	isInstalled := installer.IsInstalled()

	// Assert: Check if the package manager is correctly installed
	if !isInstalled {
		t.Error("Expected package manager to be installed, but it's not")
	}

	// Add similar tests for other package managers
}

func TestChocolateyPackageManager_Install(t *testing.T) {
	// Arrange: Create a new Chocolatey package manager
	chocoPackageManager := packagemanagers.NewChocolatey()

	// Act: Install a package
	installErr := chocoPackageManager.Install("firefox", "116.0.2")

	// Assert: Check for errors and expectations
	if installErr != nil {
		t.Errorf("Error installing package: %v", installErr)
	}

	// Cleanup: Uninstall the package (assuming successful install)
	uninstallErr := chocoPackageManager.Uninstall("firefox")
	if uninstallErr != nil {
		t.Errorf("Error uninstalling package: %v", uninstallErr)
	}
}

func TestChocolateyPackageManager_List(t *testing.T) {
	// Arrange: Create a new Chocolatey package manager
	chocoPackageManager := packagemanagers.NewChocolatey()

	// Act: List installed packages
	packages, listErr := chocoPackageManager.List()

	// Assert: Check for errors and expectations
	if listErr != nil {
		t.Errorf("Error listing packages: %v", listErr)
	}

	// Assert: Check if at least one package is listed
	if len(packages) == 0 {
		t.Errorf("Expected at least one package in the list, but got none")
	}
}
