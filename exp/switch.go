package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	var command string

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Print("Enter command: ")
	fmt.Scanln(&command)
	switch command {
	case "feed":
		fmt.Println("hunger + 5")
	case "clean":
		fmt.Println("waste false")
	case "heal":
		fmt.Println("health + 0`")
	}
}
