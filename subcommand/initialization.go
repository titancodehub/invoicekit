package subcommand

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/titancodehub/invoicekit/tui/loading"
	"github.com/titancodehub/invoicekit/tui/styles"
	"os"
	"time"
)

const (
	InitCommandName string = "init"
)

func Initialization() {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	m := loading.New(sp)
	p := tea.NewProgram(m)

	go func() {
		p.Send(loading.SendLoadingMsg("checking configuration"))
		isConfigFileExist, err := isConfigCreated()
		if err != nil {
			fmt.Println(styles.WithErrorSymbol("failed to check configuration"))
			os.Exit(1)
		}

		if isConfigFileExist {
			fmt.Println(styles.WithSuccessSymbol("configuration exist"))
			os.Exit(0)
		}

		p.Send(loading.SendLoadingMsg("configuration is not exists, generating ..."))
		if err = createConfigFile(); err != nil {
			fmt.Println(styles.WithErrorSymbol("failed to generate config"))
			time.Sleep(1 * time.Second)
			os.Exit(1)
		}

		fmt.Println(styles.WithSuccessSymbol("%s configuration is set"))
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	if _, err := p.Run(); err != nil {
		fmt.Print(styles.WithErrorSymbol("failed to execute loading command"))
		time.Sleep(1 * time.Second)
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

func createConfigFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configFile := homeDir + "/.config/invoicekit.json"
	if _, err = os.Create(configFile); err != nil {
		return err
	}

	return nil
}
