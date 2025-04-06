package main

import "fmt"

type Attacker interface {
	Attack()
}

type Warrior struct{}

func (w Warrior) Attack() {
	fmt.Println("战士用剑攻击")
}
func main() {
	var play1 Attacker = Warrior{}
	play1.Attack()
}
