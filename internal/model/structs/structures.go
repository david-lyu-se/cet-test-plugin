package structures

/* File IO  */
type Applications []Application

type Application struct {
	Name       string
	Path       string
	PluginPath string
}

// Needed to get Tea List Item to show

/* Title to display into Tea List Item */
func (a Application) Title() string { return "Appication: " + a.Name }

/* Description to display into Tea List Item */
func (a Application) Description() string { return "Path: " + a.Path }

/* How Tea List Item gets fileterd */
func (e Application) FilterValue() string {
	return e.Name
}

type ConfFile struct {
	Apps        Applications
	MonoRepoDir string
	PluginDir   string
	ThemeDir    string
	// For file picker start directory. Change this manually
	WorkingDir string
}
