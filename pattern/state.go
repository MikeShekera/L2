package pattern

import "fmt"

type CurrentSwitch interface {
	Toggle(s *LightSwitch)
}

type OnState struct{}

func (o OnState) Toggle(s *LightSwitch) {
	fmt.Println("On")
	s.SetState(&OffState{})
}

type OffState struct{}

func (o OffState) Toggle(s *LightSwitch) {
	fmt.Println("Off")
	s.SetState(&OnState{})
}

type LightSwitch struct {
	sw CurrentSwitch
}

func (l *LightSwitch) SetState(c CurrentSwitch) {
	l.sw = c
}

func (l *LightSwitch) Toggle() {
	l.sw.Toggle(l)
}
