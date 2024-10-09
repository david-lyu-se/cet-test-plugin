package structs

type Environments []Environment

type Environment struct {
	Name       string
	Path       string
	PluginPath string
}

func (e Environment) FilterValue() string {
	return e.Name
}
