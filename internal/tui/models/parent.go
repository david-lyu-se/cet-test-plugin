// package parentModel

// import (
// 	"strings"
// 	// apps "test-cet-wp-plugin/internal/tui/models/appsV2"
// 	"test-cet-wp-plugin/internal/tui/variables"

// 	"github.com/charmbracelet/bubbles/key"
// 	"github.com/charmbracelet/bubbles/list"
// 	tea "github.com/charmbracelet/bubbletea"
// )

// /* Initial variables */
// type item string

// func (i item) FilterValue() string { return string(i) }
// func (i item) Title() string       { return string(i) }
// func (i item) Description() string { return "" }

// const (
// 	enumApp      item = "Go to Application settings"
// 	enumRepo     item = "Edit mono repo directory"
// 	enumPlugin   item = "Edit Plugins Parent Directory"
// 	enumFileSync item = "Execute file sync"
// )

// /* Start of Tea Model */
// type ParentModel struct {
// 	List         list.Model
// 	hasTryDelete bool
// 	hasTryEdit   bool
// 	err          error
// 	Quitting     bool
// }

// func InitParent() (tea.Model, tea.Cmd) {

// 	var items = []list.Item{
// 		item(enumApp),
// 		item(enumRepo),
// 		item(enumPlugin),
// 		item(enumFileSync),
// 	}

// 	var parentModel = ParentModel{
// 		List:         list.New(items, list.NewDefaultDelegate(), 20, 24),
// 		hasTryDelete: false,
// 		hasTryEdit:   false,
// 		Quitting:     false,
// 	}

// 	if variables.WindowSize.Height != 0 {
// 		top, right, bottom, left := variables.DocStyle.GetMargin()
// 		parentModel.List.SetSize(variables.WindowSize.Width-left-right, variables.WindowSize.Height-top-bottom)
// 	}

// 	parentModel.List.SetSize(50, 50)
// 	parentModel.List.Title = "Start Up Menu"
// 	parentModel.List.AdditionalShortHelpKeys = func() []key.Binding {
// 		return []key.Binding{
// 			variables.Keymap.Enter,
// 			variables.Keymap.Quit,
// 		}
// 	}

// 	return parentModel, func() tea.Msg { return tea.WindowSizeMsg{} }
// }

// func (pModel ParentModel) Init() tea.Cmd {
// 	return nil
// }

// func (pModel ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

// 	var cmd tea.Cmd
// 	var cmds []tea.Cmd

// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		top, right, bottom, left := variables.DocStyle.GetMargin()
// 		variables.WindowSize = tea.WindowSizeMsg{
// 			Width:  msg.Width - left - right,
// 			Height: msg.Height - top - bottom - 1,
// 		}
// 		pModel.List.SetSize(variables.WindowSize.Width, variables.WindowSize.Height)
// 	case tea.KeyMsg:
// 		if key.Matches(msg, variables.Keymap.Enter) {
// 			var item = pModel.List.SelectedItem()
// 			return pModel.goToSubMenu(item)
// 		} else if (key.Matches(msg, variables.Keymap.Quit)) || key.Matches(msg, variables.Keymap.Quit) {
// 			pModel.Quitting = true
// 			cmds = append(cmds, tea.Quit)
// 		}
// 	default:
// 		msg = func() tea.Msg {
// 			return tea.WindowSizeMsg{
// 				Width:  variables.WindowSize.Width,
// 				Height: variables.WindowSize.Height,
// 			}
// 		}
// 	}
// 	pModel.List, cmd = pModel.List.Update(msg)
// 	cmds = append(cmds, cmd)
// 	return pModel, tea.Batch(cmds...)
// }

// func (pModel ParentModel) View() string {

// 	if pModel.Quitting {
// 		return ""
// 	}

// 	var s strings.Builder

// 	// Change to not ref
// 	if variables.AppInfo != nil {
// 		s.WriteString("Application chosen:" + variables.AppInfo.Name + "\n")
// 		s.WriteString("Application path:" + variables.AppInfo.Path + "\n")
// 		s.WriteString("testing 12345")
// 	}
// 	if variables.AppInfo == nil {
// 		s.WriteString(variables.DocStyle.Render(pModel.List.View()) + "\n")

// 	}

// 	return s.String()
// }

// /**** ---------------- ****/
// /**** Helper Functions ****/

// func (pModel ParentModel) goToSubMenu(i list.Item) (tea.Model, tea.Cmd) {
// 	switch true {
// 	case enumApp == i:
// 		// var appModel, cmd = applicationModel.InitAppModel(&variables.Conf.Apps)
// 		// variables.AppModel = &appModel
// 		// return appModel, cmd
// 		return apps.InitApps(), nil
// 	case enumRepo == i:
// 		return nil, nil
// 	case enumPlugin == i:
// 		return nil, nil
// 	case enumFileSync == i:
// 		return nil, nil
// 	default:
// 		return pModel, nil
// 	}
// }
