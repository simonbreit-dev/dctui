package state

type AppState struct {
	Items        []string
	SelectedIdx  int
	CommandMode  bool
	CommandInput string
}

func NewAppState() *AppState {
	return &AppState{
		Items: []string{"Item 1", "Item 2", "Item 3"},
	}
}
