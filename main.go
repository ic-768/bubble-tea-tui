package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	nameInput textinput.Model
	noteInput textarea.Model

	step int
	err  error
}

func initialModel() model {
	nameInput := textinput.New()
	nameInput.Placeholder = "Dwight"
	nameInput.Focus()
	nameInput.CharLimit = 156
	nameInput.Width = 20

	noteInput := textarea.New()
	noteInput.Placeholder = "Make a note..."

	return model{
		nameInput: nameInput,
		noteInput: noteInput,
		err:       nil,
		step:      1,
	}
}

func (m model) handleKey(msg tea.KeyMsg) (model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC:
		return m, tea.Quit
	case tea.KeyEnter:
		if m.step == 1 {
			m.nameInput.Blur()
			m.noteInput.Focus()
			m.step++
			return m, nil
		}
	case tea.KeyEsc:
		if m.step == 2 {
			m.step++
		}
	}

	switch m.step {
	case 1:
		m.nameInput, _ = m.nameInput.Update(msg)
	case 2:
		m.noteInput, _ = m.noteInput.Update(msg)
	}

	return m, nil
}

func (m model) Init() tea.Cmd {
	return textinput.Blink // Add the blinking command for the textarea
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)
	case error:
		return m.setError(msg), nil
	}
	return m, nil
}

func (m model) View() string {
	switch m.step {
	case 1:
		return m.viewStep1()
	case 2:
		return m.viewStep2()
	default:
		return ""
	}
}

func (m model) viewStep1() string {
	return fmt.Sprintf("What’s your name?\n\n%s", m.nameInput.View())
}

func (m model) viewStep2() string {
	return fmt.Sprintf("Tell me a story.\n\n%s\n\n%s", m.noteInput.View(), "(Esc to save)")
}

func (m model) setError(err error) model {
	m.err = err
	return m
}
