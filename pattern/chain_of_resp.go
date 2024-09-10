package pattern

type Phase interface {
	Execute(*Call)
	SetNext(Phase)
}

type Call struct {
	Index        int
	IsPhase1Done bool
	IsPhase2Done bool
}

/*type enter struct {
	next pattern.Phase
}

func (e *enter) Execute(c *pattern.Call) {
	if !c.IsPhase1Done {
		fmt.Println("Setting Index To 1")
	} else {
		fmt.Println("Phase 1 Already Done")
	}
	e.next.Execute(c)
}

func (e *enter) SetNext(next pattern.Phase) {
	e.next = next
}

type mid struct {
	next pattern.Phase
}

func (e *mid) Execute(c *pattern.Call) {
	if !c.IsPhase2Done {
		fmt.Println("Setting Index To 2")
	} else {
		fmt.Println("Phase 2 Already Done")
	}
	fmt.Println("DONE")
}

func (e *mid) SetNext(next pattern.Phase) {
	e.next = next
}

func main() {
	newMid := &mid{}

	newEnter := &enter{}
	newEnter.next = newMid

	newCall := &pattern.Call{}
	newEnter.Execute(newCall)
}*/
