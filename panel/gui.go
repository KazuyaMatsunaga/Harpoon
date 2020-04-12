package panel

import (
	"github.com/rivo/tview"

	docon "harpoon/controller"
)

type GUI struct {
	App        *tview.Application
	Pages      *tview.Pages
	ContaPanel *ContaPanel
}

func NewGUI() *GUI {
	return &GUI{
		App:        tview.NewApplication(),
		Pages:      tview.NewPages(),
		ContaPanel: newContaPanel(),
	}
}

func (g *GUI) Run() error {
	g.ContaPanel.SetConta(docon.ContaList())
	g.ContaPanel.UpdateView()

	grid := tview.NewGrid().SetRows(0, 0).AddItem(g.ContaPanel, 0, 0, 1, 1, 0, 0, true)

	g.Pages.AddAndSwitchToPage("main", grid, true)

	return g.App.SetRoot(g.Pages, true).Run()
}
