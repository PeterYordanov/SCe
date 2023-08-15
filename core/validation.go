package core

import (
	"fmt"
	"os"
)

// Includes
func ValidateIncludeExists(includePath string) (bool, error) {
	_, err := os.ReadFile(includePath)
	if err != nil {
		return false, err
	}

	fmt.Println("Found include file:", includePath)
	return true, nil
}

// Collections/Runs
func CollectionExistsInIncludes(target string, slice []string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// Packages
func IsPackageInstalled(pkgManager PackageManager, name string, version string) (bool, error) {
	packagesInstalled, err := pkgManager.List()

	if err != nil {
		return false, err
	}

	for _, value := range packagesInstalled {
		if name == value.Name && version == value.Version {
			return true, nil
		}
	}

	return false, nil
}
