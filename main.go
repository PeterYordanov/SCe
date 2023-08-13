package main

import (
	"fmt"

	"github.com/PeterYordanov/SCe/packagemanagers"
	"github.com/PeterYordanov/SCe/parser"
)

func main() {
	runbook := parser.NewRunbook("C:\\Projects\\SCe\\sample\\sample-runbook.yml")
	runbook.Parse()
	runbook.Run()

	installers := packagemanagers.NewPackageManagerInstaller(packagemanagers.Scoop)

	if !installers.IsInstalled() {
		installers.Install()
	}

	fmt.Println(installers.IsInstalled())
	fmt.Println(installers.GetVersion())

	scoop := packagemanagers.NewScoop()
	scoop.Install("1password-cli", "2.19.0")

	fmt.Println(scoop.List())

	scoop.Uninstall("1password-cli")

	/*
		if !installers.IsInstalled() {
			fmt.Println("Installing Chocolatey...")
			installers.Install()
			fmt.Println("Successfully installed Chocolatey!")
		} else {
			fmt.Println("Uninstalling Chocolatey...")
			installers.Uninstall()
			fmt.Println("Successfully uninstalled Chocolatey!")
		}*/

	fmt.Println(packagemanagers.UninstallSystemProgram("Microsoft.BingSports"))

}
