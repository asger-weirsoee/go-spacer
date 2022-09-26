package workspace

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"go.i3wm.org/i3"
	"weircon.dk/go-spacer/workspace/configuration"
)

type Output struct {
	name       string
	index      int
	workspaces []string
	width      int64
	height     int64
	x_offset   int64
	y_offset   int64
}

func GetWorkspaces(path string) []string {
	var re = regexp.MustCompile(`(?m)workspace "(\d+:.*)".*`)
	var res []string = []string{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if len(re.FindStringIndex(fileScanner.Text())) > 0 {
			kk := re.FindAllStringSubmatch(fileScanner.Text(), -1)
			println(kk[0][1])
			res = append(res, kk[0][1])
		}

	}

	return res
}

func GetOutputs(conf configuration.DefaultConfig) []Output {
	outputs, _ := i3.GetOutputs()
	res := []Output{}
	for _, o := range outputs {

		root, file, err := GetIndex(o.Name, conf.ConfRootDir)
		if err != nil {
			conf.CreateDefaultOutput(o.Name)
			root, file, err = GetIndex(o.Name, conf.ConfRootDir)
			if err != nil {
				panic(fmt.Errorf("Cannot find the index......"))
			}
		}
		index, err := strconv.Atoi(root)
		if err != nil {
			panic(fmt.Errorf("Could not deduce an int"))
		}

		res = append(res, Output{
			name:       o.Name,
			index:      index,
			workspaces: GetWorkspaces(file),
			width:      o.Rect.Width,
			height:     o.Rect.Height,
			x_offset:   o.Rect.X,
			y_offset:   o.Rect.Y,
		})
	}
	return res
}

type WorkspaceHandeler interface {
	MoveFocus(index int)
	MoveFocusedWindow(index int)
	getFocusedOutput(outputs []Output) (Output, error)
}

type I3Handeler struct {
	outputs []Output
}

func CreateHandeler(conf configuration.DefaultConfig) WorkspaceHandeler {
	if true { // Maybe support for Sway aswell?
		var i WorkspaceHandeler = I3Handeler{
			outputs: GetOutputs(conf),
		}
		return i
	}
	return nil
}

func (g I3Handeler) MoveFocus(index int) {
	out, _ := g.getFocusedOutput(g.outputs)

	i3.RunCommand(fmt.Sprintf("workspace %s", out.workspaces[index-1]))
}

func (g I3Handeler) MoveFocusedWindow(index int) {
	out, _ := g.getFocusedOutput(g.outputs)

	i3.RunCommand(fmt.Sprintf("move container to workspace %s", out.workspaces[index-1]))

}

func (g I3Handeler) getFocusedOutput(outputs []Output) (Output, error) {
	xx, yy := robotgo.GetMousePos()
	x := int64(xx)
	y := int64(yy)
	for i := 0; i < len(outputs); i++ {
		qO := outputs[i]
		if qO.x_offset == 0 && qO.y_offset == 0 {
			if qO.x_offset <= x && x <= qO.x_offset+qO.width &&
				qO.y_offset <= y && y <= qO.y_offset+qO.height {
				return qO, nil
			}
		} else if qO.x_offset == 0 {
			if qO.x_offset <= x && x <= qO.x_offset+qO.width &&
				qO.y_offset < y && y <= qO.y_offset+qO.height {
				return qO, nil
			}
		} else if qO.y_offset == 0 {
			if qO.x_offset < x && x <= qO.x_offset+qO.width &&
				qO.y_offset <= y && y <= qO.y_offset+qO.height {
				return qO, nil
			}
		} else {
			if qO.x_offset < x && x <= qO.x_offset+qO.width &&
				qO.y_offset < y && y <= qO.y_offset+qO.height {
				return qO, nil
			}
		}
	}
	return Output{}, fmt.Errorf("Could not find a focused output")
}

//func (g *i3WorkspaceGetter)

func GetIndex(name string, configRoot string) (string, string, error) {
	var re = regexp.MustCompile(`(?m)(\d*)$`)
	files, _ := ioutil.ReadDir(configRoot)
	for _, file := range files {
		if strings.HasPrefix(file.Name(), name) {
			finding := re.FindString(file.Name())
			return finding, path.Join(configRoot, file.Name()), nil
		}
	}
	return "", "", fmt.Errorf("No config for %s", name)
}
