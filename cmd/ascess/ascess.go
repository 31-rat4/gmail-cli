package ascess

type ascessHandler struct {
	state string
}

func NewAscessHandler(s string) ascessHandler {
	return ascessHandler{s}
}

func (a ascessHandler) GetConfigFromFile() string {
	return a.state
}
