// L1.21

// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// структура авто
type Car struct {
	player Player
}

func (car *Car) MusicOn() {
	car.player.TurnOn()
}

func (car *Car) MusicOff() {
	car.player.TurnOff()
}

// интерфейс плейера, который используется в авто
type Player interface {
	TurnOn()
	TurnOff()
}

// Плейер Sony со своими методами
type SonyPlayer struct {
	name string
}

func (p *SonyPlayer) VerySonyTurnOn() {
	fmt.Println(p.name, "was turned on")
}

func (p *SonyPlayer) VerySonyTurnOff() {
	fmt.Println(p.name, "was turned off")
}

// Адаптер для плейера Sony, чтобы использоватать в авто
type SonyAdapter struct {
	*SonyPlayer
}

func (player *SonyAdapter) TurnOn() {
	player.VerySonyTurnOn()
}

func (player *SonyAdapter) TurnOff() {
	player.VerySonyTurnOff()
}

func main() {
	sonyPlayer := &SonyPlayer{name: "Sony F200"}

	sonyAdapter := SonyAdapter{sonyPlayer}

	car := Car{player: &sonyAdapter}

	car.MusicOn()
	car.MusicOff()
}
