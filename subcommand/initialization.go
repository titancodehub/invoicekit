package subcommand

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/titancodehub/invoicekit/tui/initialization"
	"github.com/titancodehub/invoicekit/tui/styles"
	"os"
	"time"
)

const (
	InitCommandName string = "init"
)

func Initialization() {
	loading := spinner.New()
	loading.Spinner = spinner.Dot
	m := initialization.NewInitModel(loading)
	p := tea.NewProgram(m)

	go func() {
		p.Send(initialization.SendLoadingMsg("checking configuration"))
		time.Sleep(2 * time.Second)
		isConfigFileExist, err := isConfigCreated()
		if err != nil {
			panic(err)
		}

		if isConfigFileExist {
			p.Send(initialization.SendOutput("configuration exist"))
			return
		}

		p.Send(initialization.SendLoadingMsg("configuration is not exists, generating ..."))
		time.Sleep(2 * time.Second)
		p.Send(initialization.SendOutput(fmt.Sprintf("%s configuration is set", styles.Symbol["check"])))
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	if _, err := p.Run(); err != nil {
		fmt.Print(styles.ErrorTextStyle.Render("failed to execute initialization command"))
		os.Exit(1)
	}
}

func isConfigCreated() (bool, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}

	configFile := homeDir + "/.config/invoicekit.json"
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
