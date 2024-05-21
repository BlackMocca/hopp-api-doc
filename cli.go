package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	mets "github.com/hoppscotch/hopp-cli/methods"
	"github.com/urfave/cli"
)

// VERSION is set by `make` during the build to the most recent tag
var buildVersion = "unknown"

func main() {
	app := cli.NewApp()
	app.Name = color.HiGreenString("Hoppscotch CLI")
	app.Version = color.HiRedString(buildVersion)
	app.Usage = color.HiYellowString("Test API endpoints without the hassle")
	app.Description = color.HiBlueString("Made with <3 by Hoppscotch Team")

	var out string

	genFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Writing Directory to generate markdown and html",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "Port at which the server will open to",
			Value: 1341,
		},
		cli.BoolFlag{
			Name:  "browser, b",
			Usage: "Whether to open the browser automatically",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "gen",
			Usage: "Generate Documentation from the Hoppscotch Collection.json",
			Flags: genFlags,
			Action: func(c *cli.Context) error {
				if err := mets.GenerateDocs(c); err != nil {
					return err
				}
				return nil
			},
		},
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%s

	WE REALLY NEED YOUR FEEDBACK,

	CREATE A NEW ISSUE FOR BUGS AND FEATURE REQUESTS : < https://github.com/BlackMocca/hopp-cli >

	`, cli.AppHelpTemplate)

	err := app.Run(os.Args)
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Println(color.HiRedString("\n%s", err.Error()))
		os.Exit(1)
	}
	fmt.Fprintf(color.Output, out)
}
