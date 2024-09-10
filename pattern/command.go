package pattern

type Command interface {
	Execute()
}

type Switch interface {
	Up()
	Down()
}

type UpCommand struct {
	sw Switch
}

func (up *UpCommand) Execute() {
	up.sw.Up()
}

type DownCommand struct {
	sw Switch
}

func (up *DownCommand) Execute() {
	up.sw.Down()
}
