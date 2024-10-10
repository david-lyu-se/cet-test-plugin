package structures

/* File IO  */
type Applications []Application

type Application struct {
	Name       string
	Path       string
	PluginPath string
}

// Needed to add Application to as a Tea List Item
func (e Application) FilterValue() string {
	return e.Name
}

type ConfFile struct {
	Apps       *Applications
	WorkingDir string
	PluginDir  string
}
