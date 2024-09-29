package modelbubbletea

type EnvironmentInputs struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

type PluginTypeInputs struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}
