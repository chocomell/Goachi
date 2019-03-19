package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	cyclespeed = 3 //in seconds
	debug      = true
)

type Goachi struct {
	Age    int
	Health float64
	Hunger int
	Name   string
	Waste  bool
}

func main() {
	//start up
	// os := runtime.GOOS
	// clear(os) //blanks screen.
	clear()
	Gai := &Goachi{Name: "Tama", Age: 0, Health: 100, Waste: false, Hunger: 100}
	Gai.startup()

	go Gai.program()
	termInput()
}

func (Gai *Goachi) program() {
	//Day Cycle Loop
	for {
		//program prints:
		fmt.Println("name is:", Gai.Name)
		fmt.Println("health is:", Gai.Health)
		fmt.Println("age is:", Gai.Age)
		fmt.Println("hunger is:", Gai.Hunger)
		if Gai.Waste == true {
			fmt.Println("I am diry, consider cleaning")
		}
		if Gai.Health <= 50 {
			fmt.Println("Iam slowly dying!")
		}
		fmt.Println(">.. enter command <commands>")
		fmt.Println("==================================================================")

		//Multipliers:
		Gai.HungerDegrader()
		Gai.HealthDegrader()
		Gai.Waster()
		Gai.Age = (Gai.Age + 1)
		Sleep()
	}
}

//Sleep cycle:
func Sleep() {
	s := cyclespeed
	for i := 0; i <= s; i++ {
		time.Sleep(1 * time.Second)
	}
}

func (g *Goachi) Waster() {
	rand.Seed(time.Now().UnixNano()) //pdeudo random number generator
	r := rand.Intn(100)
	if r <= 15 {
		fmt.Println("waste generated", r)
		g.Waste = true
	}
}

func (Gai *Goachi) HealthDegrader() {
	if Gai.Waste == true {
		Gai.Health = (Gai.Health - 10) //health degration
	}
	if Gai.Health <= 0 {
		for {
			fmt.Println("GAME OVER")
			time.Sleep(9000 * time.Hour)
		}
	}
}

func (g *Goachi) HungerDegrader() {
	g.Hunger = (g.Hunger - 20) //health degration
}

func (Gai *Goachi) startup() {
	if debug != true {
		fmt.Print("Enter name: ")
		fmt.Scanln(&Gai.Name)
		fmt.Println(Gai.Name)
	}
}

func clearWin() {
	//TO DO CREATE LINUX ON AND ADD OS CHECKKING. ANDROID?

}

func clearLin() {

}

func clear() {
	var env string
	if env == "" {
		env = runtime.GOOS
	}
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	time.Sleep(5 * time.Second)
	switch env {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		fmt.Println("\033[2J") //blanks terminal unix only
	}
}

func termInput() {
	for {
		fmt.Println("start input()")
		var command string
		fmt.Scanln(&command)
		switch command {
		case "feed":
			fmt.Println("hunger + 5")
		case "clean":
			fmt.Println("waste false")
		case "heal":
			fmt.Println("health + 0")
		}
	}
}
