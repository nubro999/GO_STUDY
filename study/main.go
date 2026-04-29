package main

import (
	"fmt"
	"study/studypkg"
) 

type gasengine struct {
	fuel string
	kw   int
}

type electricengine struct {
	volt int
	kw   int
}

func (g gasengine) start() {
	fmt.Println("Gas engine started with fuel:", g.fuel)
}

func (e electricengine) start() {
	fmt.Println("Electric engine started with voltage:", e.volt)
}

type engine interface {
	start()
}

func main() {
	var eng engine

	eng = gasengine{fuel: "Petrol", kw: 150}
	eng.start()
	
	eng = electricengine{volt: 400, kw: 200}
	eng.start()

	studypkg.Study()
}
