package configuration

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

//go:embed templates/output.txt
var output string

//go:embed templates/controls.txt
var controls string

type DefaultConfig struct {
	ConfRootDir string
}

func GenDefaultConfig(basePath string) DefaultConfig {
	cc := DefaultConfig{
		ConfRootDir: basePath,
	}
	os.MkdirAll(cc.ConfRootDir, os.ModePerm)
	if _, err := os.Stat(path.Join(basePath, "controls")); errors.Is(err, os.ErrNotExist) {
		cc.CreateDefaultController()
	}
	return cc
}

func (c *DefaultConfig) CreateDefaultOutput(outputName string) {
	amountOfFiles, _ := fileCount(c.ConfRootDir)
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"add": func(i int, c int) int {
			return i + c
		},
	}
	t, _ := template.New("bob").Funcs(funcMap).Parse(output)
	data := make(map[string]interface{})
	data["name"] = outputName
	data["index"] = amountOfFiles * 10
	var tpl bytes.Buffer
	t.Execute(&tpl, data)
	thisOut := path.Join(c.ConfRootDir, fmt.Sprintf("%s_%d", outputName, data["index"]))
	os.WriteFile(thisOut, tpl.Bytes(), 0644)
}

func (c *DefaultConfig) CreateDefaultController() {
	thisOut := path.Join(c.ConfRootDir, "controls")
	os.WriteFile(thisOut, []byte(controls), 0644)
}

func fileCount(path string) (int, error) {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		if !file.IsDir() && file.Name() != "controls" {
			i++
		}
	}
	return i, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
