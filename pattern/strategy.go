package pattern

import "fmt"

type Transport interface {
	getTransport(p *Passenger) string
}

type Passenger struct {
	money int
}

type Taxi struct{}

func (t *Taxi) getTransport(p *Passenger) string {
	p.money = p.money - 2
	return "Пассажир выбрал такси"
}

type Bus struct{}

func (b *Bus) getTransport(p *Passenger) string {
	p.money = p.money - 1
	return "Пассажир выбрал автобус"
}

type registrationDesk struct {
	transp Transport
}

func (r *registrationDesk) SetTransport(transp Transport) {
	r.transp = transp
}

func (r *registrationDesk) GetPassTransport(p *Passenger) {
	if r.transp == nil {
		fmt.Println("No Transport Chosen")
		return
	}
	r.GetPassTransport(p)
}
