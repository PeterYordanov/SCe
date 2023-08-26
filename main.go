package main

import (
	"fmt"
	"os"

	"github.com/PeterYordanov/SCe/core"
	"github.com/PeterYordanov/SCe/packagemanagers"
)

func main() {

	scoop := packagemanagers.NewScoop()
	fmt.Println(scoop.List())

	list, _ := scoop.List()
	for _, value := range list {
		fmt.Println(value)
	}

	isAdmin := core.IsRunningAsAdmin()

	if !isAdmin {
		fmt.Println("You need to run this binary as an administrator.")
		os.Exit(1)
	}

	//runbook := parser.NewRunbook("C:\\Projects\\SCe\\sample\\sample-runbook.yml")
	//fmt.Println(runbook.Parse())
	//fmt.Println(runbook.Run())
}
