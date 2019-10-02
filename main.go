package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/BurntSushi/xdg"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args[1:]) == 0 {
		doc := `
Scaler [UP|DOWN]

Scaler scales config files up or down to deal with inconsistent
scaling on Wayland systems (does not depend on Wayland)
`
		fmt.Println(doc)
		os.Exit(1)
	}

	args := os.Args[1:]
	direction := args[0]

	paths := &xdg.Paths{XDGSuffix: "scaler"}
	configPath := paths.MustError(paths.ConfigFile("config.yml"))

	cfg, err := ParseConfig(configPath)
	if err != nil {
		log.Panicf("Can't parse config file: %s - %s\n", configPath, err)
	}

	templateGlob := filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "**/*.tmpl")
	templates, err := filepath.Glob(templateGlob)
	if err != nil {
		log.Panicf("Bad glob: %s\n%s\n", templateGlob, err)
	}

	for _, t := range templates {
		b, err := ioutil.ReadFile(t)
		if err != nil {
			log.Panicf("Can't open template file: %s\n%s\n", t, err)
		}

		contents := bytes.NewBuffer(b).String()
		finalPath := filepath.Join(filepath.Dir(t), Basename(t))
		output, err := os.Create(finalPath)
		if err != nil {
			log.Panicf("Can't create output file: %s\n%s\n", finalPath, err)
		}
		defer ioutil.NopCloser(output)

		appName := filepath.Dir(t[len(os.Getenv("XDG_CONFIG_HOME"))+1:])

		c, err := ConfigFor(cfg, appName, direction)
		if err != nil {
			log.Panicf("Can't find config for %s - %s\n%s\n", appName, cfg, err)
		}

		log.Printf("%s to %s with config %s\n\n", color.YellowString(t), color.YellowString(finalPath), color.RedString("%s", c))

		fileTemplate := template.Must(template.New(t).Parse(contents))
		err = fileTemplate.Execute(output, c)
		if err != nil {
			log.Panicln("Template error:", err)
		}
	}
}
