package main

import "fmt"

type TV struct{}

func (t *TV) On() {
	println("TV is on")
}

func (t *TV) Off() {
	println("TV is off")
}

type Light struct{}

func (l *Light) On() {
	println("Light is on")
}

func (l *Light) Off() {
	println("Light is off")
}

type Xbox struct{}

func (x *Xbox) On() {
	println("Xbox is on")
}

func (x *Xbox) Off() {
	println("Xbox is off")
}

type Micropone struct{}

func (m *Micropone) On() {
	println("Micropone is on")
}

func (m *Micropone) Off() {
	println("Micropone is off")
}

type Projector struct{}

func (p *Projector) On() {
	println("Projector is on")
}

func (p *Projector) Off() {
	println("Projector is off")
}

type VoiceBox struct{}

func (v *VoiceBox) On() {
	println("VoiceBox is on")
}

func (v *VoiceBox) Off() {
	println("VoiceBox is off")
}

type HomeTheaterFacade struct {
	tv        *TV
	light     Light
	xbox      Xbox
	micropone Micropone
	projector Projector
	voiceBox  VoiceBox
}

func (h *HomeTheaterFacade) Switch2KTVMode() {
	fmt.Println("Switch to KTV mode")
	h.tv.On()
	h.light.Off()
	h.micropone.On()
	h.projector.On()
	h.voiceBox.On()
}

func (h *HomeTheaterFacade) Switch2GameMode() {
	fmt.Println("Switch to game mode")
	h.tv.On()
	h.light.On()
	h.xbox.On()
	h.micropone.On()
}

func main() {
	homeTheater := new(HomeTheaterFacade)

	homeTheater.Switch2KTVMode()

	fmt.Println("------------")

	homeTheater.Switch2GameMode()
}
