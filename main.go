package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func run(c *cli.Context) error {
	level := c.String("level")
	if !stringInSlice(level, levels) {
		return fmt.Errorf("Level %q not valid. Expected one of  %s", level, levels)
	}
	experiment := c.String("experiment")
	if !stringInSlice(experiment, experiments) {
		return fmt.Errorf("Experiment %q not valid. Expected one of %s", experiment, experiments)
	}

	dir, err := ioutil.TempDir("", fmt.Sprintf("genome-%s-level_%s_", experiment, level))
	if err != nil {
		return err
	}

	for _, chr := range chromosomes {
		file := filepath.Join(dir, fmt.Sprintf("%s.json", chr))
		DownloadFile(file, fmt.Sprintf(
			"http://3dgnome.cent.uw.edu.pl/models/experiment-%s/sample-HCMConverted/%s/level-%s/coords.json",
			experiment,
			chr,
			level,
		))
	}
	log.Printf("Files written to %s\n", dir)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "fetch_genome"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "level, l",
			Value: "1",
			Usage: fmt.Sprintf("Resolution level (one of %s)", levels),
		},

		cli.StringFlag{
			Name:  "experiment, e",
			Value: "GM12878",
			Usage: fmt.Sprintf("experiment to use (one of %s)", experiments),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
