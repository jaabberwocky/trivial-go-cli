package main

import (
	"fmt"
	"os/exec"
	"simpleCLI/module1"
)

func main() {
	out, err := exec.Command("cmd", "/C", "echo", module1.ExportedString).Output()
	if err != nil {
		fmt.Println("Something went wrong!")
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
		fmt.Println("Command executed succesfully!")
	}
}
