package packagemanagers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type PackageManagerType int

const (
	Chocolatey PackageManagerType = iota
	AptGet
	Snap
	Scoop
	Winget
	Dpkg
)

var packageManagerNames = [...]string{
	"Chocolatey",
	"AptGet",
	"Snap",
	"Scoop",
	"Winget",
	"Dpkg",
}

type PackageManagerInstaller struct {
	Type PackageManagerType
}

func NewPackageManagerInstaller(packageManagerType PackageManagerType) *PackageManagerInstaller {
	return &PackageManagerInstaller{
		Type: packageManagerType,
	}
}

func (p PackageManagerInstaller) IsInstalled() bool {
	switch p.Type {
	case Chocolatey:
		cmd := exec.Command("choco", "--version")
		err := cmd.Run()
		return err == nil
	case AptGet:
		cmd := exec.Command("apt-get", "--version")
		err := cmd.Run()
		return err == nil
	case Snap:
		cmd := exec.Command("snap", "--version")
		err := cmd.Run()
		return err == nil
	case Dpkg:
		cmd := exec.Command("dpkg", "--version")
		err := cmd.Run()
		return err == nil
	case Scoop:
		cmd := exec.Command("scoop", "--version")
		err := cmd.Run()
		return err == nil
	case Winget:
		cmd := exec.Command("winget", "--version")
		err := cmd.Run()
		return err == nil
	default:
		return false
	}
}

func (p PackageManagerInstaller) GetVersion() (string, error) {
	var cmd *exec.Cmd
	var out []byte
	var err error

	switch p.Type {
	case Chocolatey:
		cmd = exec.Command("choco", "--version")
	case AptGet:
		cmd = exec.Command("apt-get", "--version")
	case Dpkg:
		cmd = exec.Command("dpkg", "--version")
	case Snap:
		cmd = exec.Command("snap", "--version")
	case Scoop:
		cmd = exec.Command("scoop", "--version")
	case Winget:
		cmd = exec.Command("winget", "--version")
	default:
		return "", fmt.Errorf("Unknown package manager")
	}

	if cmd != nil {
		out, err = cmd.Output()
		if err != nil {
			return "", err
		}
		return string(out), nil
	}

	return "", fmt.Errorf("Failed to determine version")
}

func (p PackageManagerInstaller) Install() error {

	fmt.Println("Installing package manager ", p.Type)

	switch p.Type {
	case Chocolatey:
		psScript := `Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))`

		cmd := exec.Command("powershell", "-Command", psScript)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}
	case AptGet:
		return fmt.Errorf("invalid package manager type")
	case Snap:
		return fmt.Errorf("invalid package manager type")
	case Scoop:
		psScript := `Set-ExecutionPolicy RemoteSigned -scope CurrentUser; iex "& {$(irm get.scoop.sh)} -RunAsAdmin"`

		cmd := exec.Command("powershell", "-Command", psScript)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}
	case Winget:
		return fmt.Errorf("invalid package manager type")
	default:
		return fmt.Errorf("invalid package manager type")
	}

	return nil
}

func (p PackageManagerInstaller) Uninstall() error {

	fmt.Println("Installing package manager ", p.Type)

	switch p.Type {
	case Chocolatey:
		// Get Install Path
		chocoInstallDir := os.Getenv("ChocolateyInstall")
		if chocoInstallDir == "" {
			fmt.Println("Chocolatey installation directory not found.")
			return fmt.Errorf("Chocolatey installation directory not found")
		}

		chocoInstallPath := strings.TrimSpace(chocoInstallDir)

		// Check if it exists
		_, err := os.Stat(chocoInstallPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Chocolatey installation directory does not exist.")
			} else {
				fmt.Printf("Error checking Chocolatey installation directory: %v\n", err)
			}
			return fmt.Errorf("Chocolatey installation directory does not exist")
		}

		// Remove Chocolatey
		err = os.RemoveAll(chocoInstallPath)
		if err != nil {
			fmt.Printf("Error deleting Chocolatey installation directory: %v\n", err)
			return fmt.Errorf("Error deleting Chocolatey installation directory")
		}

		fmt.Println("Chocolatey installation directory deleted successfully.")
	case AptGet:
		return fmt.Errorf("invalid package manager type")
	case Snap:
		return fmt.Errorf("invalid package manager type")
	case Scoop:
		return fmt.Errorf("invalid package manager type")
	case Winget:
		return fmt.Errorf("invalid package manager type")
	default:
		return fmt.Errorf("invalid package manager type")
	}

	return nil
}
