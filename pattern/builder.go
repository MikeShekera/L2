package pattern

type IBuilder interface {
	SetDoors(doorsCount int)
	SetFloor(floorCount int)
	SetColor(color string)
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "small" {
		return NewSmallBuilder()
	}
	return nil
}

type House struct {
	doors int
	floor int
	color string
}

type smallHouseBuilder struct {
	doors int
	floor int
	color string
}

func (s *smallHouseBuilder) SetDoors(doors int) {
	s.doors = doors
}

func (s *smallHouseBuilder) SetFloor(floorCount int) {
	s.floor = floorCount
}

func (s *smallHouseBuilder) SetColor(newColor string) {
	s.color = newColor
}

type mediumHouseBuilder struct {
	doors int
	floor int
	color string
}

func NewSmallBuilder() *smallHouseBuilder {
	return &smallHouseBuilder{}
}

func NewMediumBuilder() *mediumHouseBuilder {
	return &mediumHouseBuilder{}
}
