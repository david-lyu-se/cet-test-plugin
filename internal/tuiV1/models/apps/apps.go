package applicationModel

import (
	"errors"
	structures "test-cet-wp-plugin/internal/model/structs"
	"test-cet-wp-plugin/internal/tui/variables"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	List         list.Model
	fileModel    *tea.Model
	hasTryDelete bool
	hasTryEdit   bool
	err          error
}

// Todo not sure if this is needed
type updateListMsg struct {
	Item structures.Application
}

func InitAppModel(apps *structures.Applications) (tea.Model, tea.Cmd) {

	if variables.AppModel != nil {
		return *variables.AppModel, func() tea.Msg {
			return variables.InitAppModel{}
		}
	}

	var items = make([]list.Item, len(*apps))
	for i, proj := range *apps {
		items[i] = list.Item(proj)
	}

	var app = AppModel{
		List:         list.New(items, list.NewDefaultDelegate(), 20, 24),
		hasTryDelete: false,
		hasTryEdit:   false,
	}

	app.List.SetSize(50, 50)
	app.List.Title = "Applications List"
	app.List.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			variables.Keymap.Create,
			variables.Keymap.Enter,
			variables.Keymap.Back,
		}
	}

	return app, func() tea.Msg {
		return variables.InitAppModel{}
	}
}

func (appModel AppModel) Init() tea.Cmd {
	return nil
}

func (appModel AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case variables.InitAppModel:
		appModel.List.SetSize(variables.WindowSize.Width, variables.WindowSize.Height)
	// custom msg for updating Item List after creatihng new Application
	case updateListMsg:
		cmds = append(cmds, appModel.List.InsertItem(len(variables.Conf.Apps)-1, msg.Item))
		appModel.List, cmd = appModel.List.Update(variables.Conf.Apps)
		cmds = append(cmds, cmd)
		return appModel, tea.Batch(cmds...)
	// key inputs
	case tea.KeyMsg:
		return appModel.handleKeyInputs(msg)
		//cant fallthroug
	}
	appModel.List, cmd = appModel.List.Update(msg)
	cmds = append(cmds, cmd)
	return appModel, tea.Batch(cmds...)
}

func (appModel AppModel) View() string {
	if appModel.hasTryDelete {
		return variables.AlertStyle("Cannot delete please go into file and delete") + "\n" + variables.DocStyle.Render(appModel.List.View())
	}

	if appModel.err != nil {
		return variables.DocStyle.Render(appModel.err.Error())
	}
	return variables.DocStyle.Render(appModel.List.View()) + "\n"
}

/**** ---------------- ****/
/**** Helper Functions ****/

func (appModel AppModel) handleKeyInputs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if !key.Matches(msg, variables.Keymap.Quit) {
		appModel.hasTryDelete = false
		appModel.hasTryEdit = false
	}

	//clears out no delete/edit msg in view
	if !appModel.hasTryDelete {
		appModel.hasTryDelete = false
	}
	if !appModel.hasTryEdit {
		appModel.hasTryEdit = false
	}

	switch true {
	case key.Matches(msg, variables.Keymap.Create):
		var fileModel tea.Model
		fileModel, cmd = InitFileModel(&appModel)
		appModel.fileModel = &fileModel
		return fileModel, cmd
	case key.Matches(msg, variables.Keymap.Delete):
		appModel.hasTryDelete = true
	case key.Matches(msg, variables.Keymap.Enter):
		//should not error out since we know for a fact this list item is application
		// variables.AppInfo, _ = appModel.List.SelectedItem().(structures.Application)
		var (
			appRef structures.Application
			ok     bool
		)
		appRef, ok = appModel.List.SelectedItem().(structures.Application)
		if !ok {
			appModel.err = errors.New("Could not type cast")
			return appModel, func() tea.Msg { return appModel.err }
		}
		variables.AppInfo = &appRef
		fallthrough
	case key.Matches(msg, variables.Keymap.Quit):
		return *variables.ParentModel, func() tea.Msg {
			return tea.WindowSizeMsg{
				Width:  variables.WindowSize.Width,
				Height: variables.WindowSize.Height,
			}
		}
	default:
		appModel.List, cmd = appModel.List.Update(msg)
	}
	return appModel, cmd
}

func returnModel(model tea.Model, cmds []tea.Cmd) (tea.Model, tea.Cmd) {
	variables.AppModel = &model
	return model, tea.Batch(cmds...)
}
