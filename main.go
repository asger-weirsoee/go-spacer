package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"weircon.dk/go-spacer/workspace"
	"weircon.dk/go-spacer/workspace/configuration"
)

//default_home = Path.home().joinpath('.config').joinpath('sway').joinpath('outputs')

var (
	usrHome, _  = os.UserHomeDir()
	defaultHome = path.Join(usrHome, ".config", "i3", "outputs")

	basepath   string
	index      int
	shift      bool
	shiftAndGo bool
)

func main() {
	flag.StringVar(&basepath, "confbase", defaultHome, "The root folder of where the configurations are going to be put.")
	flag.IntVar(&index, "index", 0, "The workspace that we will move to")
	flag.BoolVar(&shift, "shift", false, "If you wanna move the active window to the workspace")
	flag.BoolVar(&shiftAndGo, "go", false, "In combination with shift, moves the active workspace along with moving the active window")
	flag.Parse()
	if index < 1 {
		panic(fmt.Errorf("index is required, and needs to be more than one"))
	}

	baseConf := configuration.GenDefaultConfig(basepath)

	handeler := workspace.CreateHandeler(baseConf)

	if shift {
		handeler.MoveFocusedWindow(index)
		if !shiftAndGo {
			return
		}
	}
	handeler.MoveFocus(index)
}
