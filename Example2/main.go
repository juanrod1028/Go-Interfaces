package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	kickBall()
}

type FootbalPlayerTipe1 struct {
	stamina int
	power   int
}

func (f FootbalPlayerTipe1) kickBall() {
	shot := f.stamina + f.power
	fmt.Println("tipe 1 Shot", shot)
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

func main() {
	team := make([]Player, 11)

	for i := 0; i < len(team)-1; i++ {
		team[i] = FootbalPlayerTipe1{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}

	team[len(team)-1] = FootbalPlayerTipe2{
		stamina: rand.Intn(10),
		power:   rand.Intn(10),
		speed:   rand.Intn(10),
	}
	for i := 0; i < len(team); i++ {
		team[i].kickBall()
	}
}
