package pattern

type IEngine interface {
	SetName(name string)
	GetName() string
	SetPower(hp int)
	GetPower() int
}

/*type ladaEngine struct {
	hp   int
	name string
}

func (l *ladaEngine) SetName(name string) {
	l.name = name
}

func (l *ladaEngine) GetName() string {
	return l.name
}

func (l *ladaEngine) SetPower(hp int) {
	l.hp = hp
}

func (l *ladaEngine) GetPower() int {
	return l.hp
}

type porcheEngine struct {
	hp   int
	name string
}

func (p porcheEngine) SetName(name string) {
	p.name = name
}

func (p porcheEngine) GetName() string {
	return p.name
}

func (p porcheEngine) SetPower(hp int) {
	p.hp = hp
}

func (p porcheEngine) GetPower() int {
	return p.hp
}

func main() {
	lada := ladaEngine{}
	porche := porcheEngine{}
	lada.SetName("LADA")
	lada.SetPower(100)
	porche.SetName("Porche")
	porche.SetPower(650)
}
*/
