package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	cyclespeed = 3 //in seconds
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

	Gai := &Goachi{Name: "Tama", Age: 0, Health: 100, Waste: false, Hunger: 100}
	//First user stuff:
	// fmt.Print("Enter name: ")
	// fmt.Scanln(&Gai.Name)

	// fmt.Println(Gai.Name)
	fmt.Println("\033[2J") //blanks terminal unix only

	for { //Day Cycle Loop
		Sleep()

		//prgram prints:
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
		fmt.Println("==================================================================")

		var command string
		fmt.Print("Enter command: ")
		fmt.Scanln(&command)
		switch command {
		case "feed":
			fmt.Println("hunger + 5")
		case "clean":
			fmt.Println("waste false")
		case "heal":
			fmt.Println("health + 0")
		}
		//TO DO MAKE APP CONTINUE ON NO INPUT

		//Multipliers:
		Gai.HungerDegrader()
		Gai.HealthDegrader()
		Gai.Waster()
		Gai.Age = (Gai.Age + 1)
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

func (g *Goachi) HealthDegrader() {
	if g.Hunger <= 0 {
		for {
			fmt.Println("GAME OVER")
			time.Sleep(1000 * time.Second)
		}
	}
	if g.Waste == true {
		g.Health = (g.Health - 10) //health degration
	}
}

func (g *Goachi) HungerDegrader() {
	g.Hunger = (g.Hunger - 20) //health degration
}

func clearWin() {
	//TO DO CREATE LINUX ON AND ADD OS CHECKKING. ANDROID?
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
