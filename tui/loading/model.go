package loading

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/titancodehub/invoicekit/common"
	"github.com/titancodehub/invoicekit/tui/styles"
)

type InitModel struct {
	showLoading   bool
	loadingString string
	sp            spinner.Model
	output        string
}

func New(loading spinner.Model) InitModel {
	return InitModel{
		showLoading:   true,
		loadingString: "",
		sp:            loading,
	}
}

func (m InitModel) Init() tea.Cmd {
	return m.sp.Tick
}

func (m InitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return common.RegisterKey(m, msg)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.sp, cmd = m.sp.Update(msg)
		return m, cmd
	case SendLoadingMsg:
		m.showLoading = true
		m.loadingString = msg.String()
		return m, nil
	}

	return m, nil
}

func (m InitModel) View() string {
	if m.showLoading {
		return styles.InfoTextStyle.Render(fmt.Sprintf("%s %s", m.sp.View(), m.loadingString))
	}

	return ""
}
