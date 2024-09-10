package pattern

import "fmt"

type Visitor interface {
	Visit1(a1 *Action1)
	Visit2(a2 *Action2)
}

type Do interface {
	Accept(v Visitor)
}

type Action1 struct {
	index int
}

func (a *Action1) Accept(v Visitor) {
	v.Visit1(a)
}

type Action2 struct {
	index int
}

func (a *Action2) Accept(v Visitor) {
	v.Visit2(a)
}

type Person1 struct{}

func (p *Person1) Visit1(a *Action1) {
	a.index = 1
	fmt.Println(a.index)
}

func (p *Person1) Visit2(a *Action2) {
	a.index = 2
	fmt.Println(a.index)
}

type Person2 struct{}

func (p *Person2) Visit1(a *Action1) {
	a.index = 3
	fmt.Println(a.index)
}

func (p *Person2) Visit2(a *Action2) {
	a.index = 4
	fmt.Println(a.index)
}
