package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	kickBall()
	name()
}

type FootbalPlayerTipe1 struct {
	stamina int
	power   int
}

func (f FootbalPlayerTipe1) kickBall() {
	shot := f.stamina + f.power
	fmt.Println("tipe 1 Shot", shot)
}
func (f FootbalPlayerTipe1) name() {
	fmt.Print("jugador 1")
}

type FootbalPlayerTipe2 struct {
	stamina int
	power   int
	speed   int
}

func (f FootbalPlayerTipe2) kickBall() {
	shot := f.stamina + f.power + f.speed
	fmt.Println("Tipe 2 Shot", shot)
}
func (f FootbalPlayerTipe2) name() {
	fmt.Print("jugador 2")
}

func main() {
	player1 := FootbalPlayerTipe1{
		stamina: rand.Intn(10),
		power:   rand.Intn(10),
	}
	player1.kickBall()

	player2 := FootbalPlayerTipe2{
		stamina: rand.Intn(10),
		power:   rand.Intn(10),
		speed:   rand.Intn(10),
	}
	player2.kickBall()

	team := make([]Player, 4)

	for i := 0; i < len(team)-2; i++ {
		//i can change here the type of player whithout affect the logic
		team[i] = FootbalPlayerTipe1{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}

	for i := 2; i < len(team); i++ {
		//i can change here the type of player whithout affect the logic
		team[i] = FootbalPlayerTipe2{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
			speed:   rand.Intn(10),
		}
	}

	for i := 0; i < len(team); i++ {
		team[i].kickBall()
	}
}
