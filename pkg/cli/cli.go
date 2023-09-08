package cli

import (
	"github.com/urfave/cli/v2"
)

func StartCLI() *cli.App {

	app := &cli.App{
		Name:   "BookBrowser",
		Usage:  "Host your ebooks!",
		Action: RunServer,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:      "book-dir",
				Usage:     "Specifies the path to look for ebooks in.",
				Required:  false,
				Hidden:    false,
				Value:     ".",
				EnvVars:   []string{"BOOK_BROWSER_BOOK_DIR"},
				TakesFile: false,
			},
			&cli.PathFlag{
				Name:     "tmp-dir",
				Aliases:  []string{"t"},
				Usage:    "the directory to store temp files such as cover thumbnails (created on start, deleted on exit unless already exists)",
				Required: false,
				Value:    "/tmp/bookbrowser",
				EnvVars:  []string{"BOOK_BROWSER_TMP_DIR"},
			},
			&cli.StringFlag{
				Name:     "address",
				Aliases:  []string{"a"},
				Usage:    "Specify the address the server will listen on.",
				Required: false,
				Value:    "0.0.0.0",
			},
			&cli.IntFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Specify the port to listen in.",
				Required: false,
				Value:    8090,
			},
			&cli.BoolFlag{
				Name:  "clean",
				Usage: "Whether or not we want to clean up the temporary directory, if it already exists.",
				Value: false,
			},
		},
	}

	return app

}
