package server

func NewAgent() *Agent {
	a := &Agent{}
	a.Storage = NewStorage()
	return a
}

func (a *Agent) In(address string) bool {
	return a.Storage.In(address)
}

func (a *Agent) Remove(address string) {
	a.Storage.Remove(address)
}

func (a *Agent) Add(address string) {
	a.Storage.Add(address)
}
