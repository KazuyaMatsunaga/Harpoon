package panel

import (
	"github.com/rivo/tview"

	"github.com/docker/docker/api/types"
)

type ContaPanel struct {
	contas []types.Container
	*tview.Table
}

func newContaPanel() *ContaPanel {
	p := &ContaPanel{
		Table: tview.NewTable(),
	}

	p.SetBorder(true).SetTitle(" Container List").SetTitleAlign(tview.AlignLeft)

	p.SetSelectable(true, false)

	return p
}

func (c *ContaPanel) SetConta(contas []types.Container) {
	c.contas = contas
}

func (c *ContaPanel) UpdateView() {
	table := c.Clear()

	for i, ci := range c.contas {
		cell := tview.NewTableCell(ci.Names[0])
		table.SetCell(i, 0, cell)
	}
}
