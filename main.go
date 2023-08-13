package main

import (
	"fmt"

	"github.com/PeterYordanov/SCe/parser"
)

func main() {
	runbook := parser.NewRunbook("C:\\Projects\\SCe\\sample\\sample-runbook.yml")
	fmt.Println(runbook.Parse())
	//fmt.Println(runbook.Run())
}
