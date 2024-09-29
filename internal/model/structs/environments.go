package structs

type Environments []Environment

type Environment struct {
	name       string
	path       string
	pluginPath string
}
