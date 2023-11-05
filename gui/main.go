package gui

import (
	"pwd-entropy/pkg/entropy"
	"regexp"

	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
)

var (
	shouldFocus   = true
	mapping       map[*regexp.Regexp]float64
	width         = 320
	height        = 240
	title         = "Password entropy calculator"
	inputPassword string
	state         State
)

type State struct {
	length  int
	entropy int
}

// Run runs the password entropy calculator GUI
func Run() {
	state = State{
		length:  0,
		entropy: 0,
	}
	mapping = entropy.NewNamespaceToEntropy(entropy.DefaultRules)
	wnd := g.NewMasterWindow(title, width, height, 0)
	wnd.Run(loop)
}

func loop() {
	g.SingleWindow().Layout(
		g.Custom(func() {
			if shouldFocus {
				shouldFocus = false
				g.SetKeyboardFocusHere()
			}
		}),
		g.Row(
			giu.InputText(&inputPassword).Size(300).OnChange(func() {
				state.length = len(inputPassword)
				state.entropy = int(entropy.GetEntropy(mapping, inputPassword))
			}).Hint("type a password"),
		),
		g.Row(
			giu.Labelf("Length : %d", state.length),
		),
		g.Row(
			giu.Labelf("Entropy : %d bits", state.entropy),
		),
	)
}
